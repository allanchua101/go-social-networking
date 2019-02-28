package messaging

import (
	"github.com/streadway/amqp"
)

// BuildChannel method is used for building channels
// that could be used to send consume messages from
// the queues.
func BuildChannel(conn *amqp.Connection) *amqp.Channel {
	if conn == nil {
		return nil
	}
		
	chn, err := conn.Channel()

	if err != nil {
		return nil
	}

	return chn
}