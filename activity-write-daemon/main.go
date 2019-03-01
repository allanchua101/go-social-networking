package main

import (
	"os"
	"log"

	"encoding/json"

	"activity-write-daemon/messaging"
	"activity-write-daemon/models"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatal(msg)
	}
}

// main is the application's composition root.
func main() {
	log.Println("GO Social Write Daemon Starting....")

	conn := messaging.ReliableConnectionBuilder()
	chn := messaging.ReliableChannelBuilder(conn)
	err := chn.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)

	failOnError(err, "Write daemon cannot configure QOS.")

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
			var activity *models.Activity
			activityData := []byte(d.Body)
			err := json.Unmarshal(activityData, &activity)

			if err != nil {
				log.Printf("Error %v\n", err)
				log.Printf("This message cannot be unmarshalled: %s\n", d.Body)
				break
			}

			log.Printf("Activity %s is being processed..\n", activity.ID)
			d.Ack(false)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}