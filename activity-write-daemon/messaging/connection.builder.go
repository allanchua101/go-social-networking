package messaging

import (
	"log"
	"time"

	"github.com/streadway/amqp"
)


func buildConnection(mqConnStr string) (*amqp.Connection) {
	conn1, err1 := amqp.Dial(mqConnStr)

	if err1 == nil {
		return conn1
	}

	return nil
}

// ReliableConnectionBuilder method is used for retrieving
// a connection object to the target queue. This method
// performs retry every 5 seconds if the connection cannot 
// be established 
func reliableConnectionBuilder(mqConnStr string) *amqp.Connection {
	for {
		conn := buildConnection(mqConnStr)

		if conn == nil {
			log.Println("Write daemon cannot connect to target queue. Retrying in 5 seconds	....")
			time.Sleep(5 * time.Second)
		} else {
			return conn
		}
	}
}