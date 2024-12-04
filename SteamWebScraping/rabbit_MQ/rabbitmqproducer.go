package rabbit_MQ

import (
	"log"

	"github.com/streadway/amqp"
)

func SendMessage(service *RabbitMQService, exchangeName, routingKey, message string) error {
	err := service.Channel.ExchangeDeclare(
		exchangeName,
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Printf("Failed to declare exchange: %v", err)
		return err
	}

	err = service.Channel.Publish(
		exchangeName,
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType:  "text/plain",
			Body:         []byte(message),
			DeliveryMode: amqp.Persistent,
		},
	)
	if err != nil {
		log.Printf("Failed to publish message: %v", err)
		return err
	}

	log.Printf("Message published: %s", message)
	return nil
}
