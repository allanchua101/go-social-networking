package messaging

import (
	"os"

	"github.com/streadway/amqp"
)

// BuildConnection method is used for retrieving a connection 
// to master queue, the method would return the slave queue
// if master queue is not detected.
func BuildConnection() (*amqp.Connection) {
	conn1, err1 := amqp.Dial(os.Getenv("WRITE_MQ_MASTER_CONN_STR"))

	if err1 == nil {
		return conn1
	}

	conn2, err2 := amqp.Dial(os.Getenv("WRITE_MQ_SLAVE_CONN_STR"))

	if err2 == nil {
		return conn2
	}

	return nil
}