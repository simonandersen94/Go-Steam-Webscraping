package rabbit_MQ

import (
	"github.com/streadway/amqp"
	"log"
)

type RabbitMQService struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
}

func NewRabbitMQService(uri string, ClientProvidedName string) (*RabbitMQService, error) {
	conn, err := amqp.DialConfig(uri, amqp.Config{
		Properties: amqp.Table{"connection_name": ClientProvidedName},
	})
	if err != nil {
		log.Println("Failed to connect to RabbitMQ: %v", err)
		return nil, err
	}
	channel, err := conn.Channel()
	if err != nil {
		log.Println("Failed to open a channel: %v", err)
		conn.Close()
		return nil, err
	}

	return &RabbitMQService{Connection: conn, Channel: channel}, nil
}

func (r *RabbitMQService) Close() {
	if r.Channel != nil {
		r.Channel.Close()
	}
	if r.Connection != nil {
		r.Connection.Close()
	}
}
