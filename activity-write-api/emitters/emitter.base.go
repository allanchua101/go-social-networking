package emitters

import (
	"os"
	"errors"
	"github.com/streadway/amqp"
)

func bootWriteMQs() (*amqp.Connection, *amqp.Connection, bool) {
	conn1, err1 := amqp.Dial(os.Getenv("WRITE_MQ_MASTER_CONN_STR"))
	conn2, err2 := amqp.Dial(os.Getenv("WRITE_MQ_SLAVE_CONN_STR"))
	isAnyMQOpened :=  err1 == nil || err2 == nil

	return conn1, conn2, isAnyMQOpened
}

func emitEvent(conn *amqp.Connection, queueName string, jsonString string) bool {
	if(conn == nil) {
		return  false
	}

	defer conn.Close()
	ch, err := conn.Channel()
	defer ch.Close()

	if(err != nil) {
		return false
	}

	publishErr := ch.Publish(
		"", 				// Empty Exchange
		queueName,	// Worker Queue
		false, 			// No Mandatory Queue Destination
		false,			// Server can chill and wait to queue
		amqp.Publishing {
			DeliveryMode: amqp.Persistent,
			ContentType: "text/plain",
			Body: []byte(jsonString),
	})

	return publishErr == nil
}

// PublishEvent method is used for emitting events to
// redundant queue pair that will ensure message is always
// released and would reach background daemons for registering 
// event store.
func PublishEvent(queueName string, content string) error {
	master, slave, isAnyExisting := bootWriteMQs()

	if !isAnyExisting {
		return errors.New("No queue can be accessed to emit event.")
	}

	isEmittedViaMaster := emitEvent(master, queueName, content)

	if(isEmittedViaMaster) {
		return nil
	}

	isEmittedViaSlave := emitEvent(slave, queueName, content)

	if isEmittedViaSlave {
		return nil
	}

	return errors.New("Cannot publish to both master and slave queues.")
}