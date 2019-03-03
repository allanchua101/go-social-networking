package messaging

import (
	"github.com/streadway/amqp"
)

// ConsumeQueue method is used to consume messages
// from a queue.
func ConsumeQueue(mqConnStr, queueName string) (<-chan amqp.Delivery, error) {
	conn := reliableConnectionBuilder(mqConnStr)
	chn := reliableChannelBuilder(conn)

	return chn.Consume(
		queueName, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
}