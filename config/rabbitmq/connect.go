package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

var (
	exchange_name = "add_to_cart_topic"
	exchange_type = "topic"

	routing_key = "cart.item.added"
)

func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func RabbitConnection() *amqp.Channel {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	FailOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	FailOnError(err, "Failed to open a channel")

	return ch

}
