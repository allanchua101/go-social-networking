package main

import (
	"os"
	"log"

	"my-feed-daemon/messaging"
	"my-feed-daemon/parsers"
)

func main() {
	log.Printf("My feed daemon starting....")

	mqConnStr := os.Getenv("READ_MQ_CONN_STR")
	queueName := os.Getenv("MY_FEED_PROJECTOR_QUEUE_NAME")
	msgs, err := messaging.ConsumeQueue(mqConnStr, queueName)

	if err != nil  {
		log.Fatal("My feed daemon cannot consume messages")
	}

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			activity, err := parsers.ParseActivityJson(d.Body)

			if err != nil {
				log.Printf("Error %v\n", err)
				log.Printf("This event cannot be unmarshalled: %s\n", d.Body)
				break
			}

			log.Printf("Processing %s", activity.ID)

			d.Ack(false)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}