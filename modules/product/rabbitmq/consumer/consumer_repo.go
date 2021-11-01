package consumer

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

var (
	exchange_name = "add_to_cart_topic"
	exchange_type = "topic"

	routing_key = "cart.item.added"
)

type RabbitRepository struct {
	rabbitChannel *amqp.Channel
}

func NewRabbitmqRepository(rabbit *amqp.Channel) *RabbitRepository {
	return &RabbitRepository{
		rabbit,
	}
}

func (rabbit *RabbitRepository) createExchange() error {
	err := rabbit.rabbitChannel.ExchangeDeclare(
		exchange_name,
		exchange_type,
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Printf("Failed to declare an exchange: %s", err)
	}

	log.Printf("[*] Exchange created %s", exchange_name)
	return err
}

func (rabbit *RabbitRepository) Publish() {
	rabbit.createExchange()

	body := "test"

	err := rabbit.rabbitChannel.Publish(
		exchange_name,
		routing_key,
		false,
		false,
		amqp.Publishing{
			ContentType: "json",
			Body:        []byte(body)},
	)

	if err != nil {
		log.Printf("Failed to publish message: %s", err)
	}

	log.Printf("[x] Sent %s", body)

	defer rabbit.rabbitChannel.Close()
}
