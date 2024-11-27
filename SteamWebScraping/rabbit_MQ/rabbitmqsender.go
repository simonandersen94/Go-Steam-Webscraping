package rabbit_MQ

import (
	"github.com/streadway/amqp"
	"log"
)

func SendMessage(service *RabbitMQService, exchangeName, routingKey, message string) error {
	err := service.Channel.Publish(
		exchangeName,
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	if err != nil {
		log.Printf("Failed to publish message: %v", err)
		return err
	}
	log.Printf("Message published: %s", message)
	return nil
}
