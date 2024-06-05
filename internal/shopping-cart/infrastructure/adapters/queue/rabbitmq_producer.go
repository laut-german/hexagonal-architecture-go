package queue

import (
	"context"
	"hexagonal-architecture-go/internal/shopping-cart/domain/ports/queue"

	"github.com/rabbitmq/amqp091-go"
)

type RabbitMQProducer struct {
	channel   *amqp091.Channel
	queueName string
}

func NewRabbitMQProducer(channel *amqp091.Channel, queueName string) queue.QueueProducer {
	return &RabbitMQProducer{
		channel:   channel,
		queueName: queueName,
	}
}

func (p *RabbitMQProducer) Publish(ctx context.Context, message []byte) error {
	return p.channel.PublishWithContext(ctx,
		"",
		p.queueName,
		false,
		false,
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        message,
		},
	)
}
