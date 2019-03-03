package messaging

import (
	"errors"
	"github.com/streadway/amqp"
)

func bootWriteMQs(conStr1, conStr2 string) (*amqp.Connection, *amqp.Connection, bool) {
	conn1, err1 := amqp.Dial(conStr1)
	conn2, err2 := amqp.Dial(conStr2)
	isAnyMQOpened :=  err1 == nil || err2 == nil

	return conn1, conn2, isAnyMQOpened
}

func emitEvent(conn *amqp.Connection, exchangeName string, jsonString string) bool {
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
		exchangeName, 	// Empty Exchange
		"",																		// Worker Queue
		false, 																// No Mandatory Queue Destination
		false,																// Server can chill and wait to queue
		amqp.Publishing {
			DeliveryMode: amqp.Persistent,
			ContentType: "text/plain",
			Body: []byte(jsonString),
	})

	return publishErr == nil
}

// PublishEvent method is used for emitting events to
// redundant exchange pair that will ensure message delivery
func PublishEvent(conStr1, conStr2, exchangeName, content string) error {
	master, slave, isAnyExisting := bootWriteMQs(conStr1, conStr2)

	if !isAnyExisting {
		return errors.New("No queue can be accessed to emit event.")
	}

	isEmittedViaMaster := emitEvent(master, exchangeName, content)

	if(isEmittedViaMaster) {
		return nil
	}

	isEmittedViaSlave := emitEvent(slave, exchangeName, content)

	if isEmittedViaSlave {
		return nil
	}

	return errors.New("Cannot publish to both master and slave exchanges.")
}