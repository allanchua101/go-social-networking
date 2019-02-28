package main

import (
	"os"
	"log"
	"time"

	"activity-write-daemon/messaging"

	"github.com/streadway/amqp"
)

func chill() {
	time.Sleep(5 * time.Second)
}

func reliableConnectionBuilder() *amqp.Connection {
	for {
		conn := messaging.BuildConnection()

		if conn == nil {
			log.Println("Write daemon cannot connect to either master or slave queue. Retrying in 5 seconds	....")
			chill()
		} else {
			return conn
		}
	}
}

func reliableChannelBuilder(conn *amqp.Connection) *amqp.Channel {
	for {
		chn := messaging.BuildChannel(conn)

		if chn == nil {
			log.Println("Write daemon cannot open a channel. Retrying in 5 seconds....")
			chill()
		} else {
			return chn
		}
	}
}

// main is the application's composition root.
func main() {
	log.Println("GO Social Write Daemon Starting....")

	conn := reliableConnectionBuilder()
	chn := reliableChannelBuilder(conn)
	err := chn.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)

	if err != nil {
		log.Fatal("Write daemon cannot configure QOS.")
	}

	msgs, err := chn.Consume(
		os.Getenv("WRITE_API_QUEUE_NAME"), // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	if err != nil {
		log.Fatal("Write daemon cannot consume messages")
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s\n", d.Body)
			d.Ack(false)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}