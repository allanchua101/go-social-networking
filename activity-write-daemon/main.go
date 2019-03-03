package main

import (
	"os"
	"log"
	"encoding/json"

	"activity-write-daemon/messaging"
	"activity-write-daemon/persistence"
	"activity-write-daemon/parsers"
	"activity-write-daemon/models"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatal(msg)
	}
}

func storeEvent(activity *models.Activity) {
	log.Printf("Event %s is being processed..\n", activity.ID)
	isStored := persistence.PushEvent(activity)

	if isStored {
		log.Printf("Event stored?: %s\n", activity.ID)
	}
}

func publishProjectionMsg(activity *models.Activity) error {
	mqStr1 := os.Getenv("READ_MQ_CONN_STR")
	mqStr2 := os.Getenv("READ_MQ_SLAVE_CONN_STR")
	exchangeName := os.Getenv("ACTIVITY_EXCHANGE_NAME")
	serializedData, serializeErr := json.Marshal(activity)

	if serializeErr != nil {
		return serializeErr
	}
	
	err := messaging.PublishEvent(mqStr1, mqStr2, exchangeName, string(serializedData))

	return err
}

// main is the application's composition root.
func main() {
	log.Println("GO Social Write Daemon Starting....")

	conn := messaging.ReliableConnectionBuilder()
	chn := messaging.ReliableChannelBuilder(conn)
	msgs, err := chn.Consume(
		os.Getenv("WRITE_API_QUEUE_NAME"), // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	failOnError(err, "Write daemon cannot consume messages")

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			activity, err := parsers.ParseActivityJson(d.Body)

			if err != nil {
				log.Printf("Error %v\n", err)
				log.Printf("This event cannot be unmarshalled: %s\n", d.Body)
				break
			}

			storeEvent(activity)
			err = publishProjectionMsg(activity)

			if err != nil {
				log.Printf("Publishing to projection exchange failed: %s \n", activity.ID)
				break;
			}

			d.Ack(false)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}