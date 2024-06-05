package queue

import (
	"context"
	"encoding/json"
	"hexagonal-architecture-go/internal/shopping-cart/application/commands"
	"hexagonal-architecture-go/internal/shopping-cart/domain/ports/queue"
	"hexagonal-architecture-go/internal/shopping-cart/domain/ports/repositories"
	"log"

	"github.com/rabbitmq/amqp091-go"
)

type RabbitMQConsumer struct {
	channel           *amqp091.Channel
	queueName         string
	addItemHandler    *commands.AddItemToCartHandler
	removeItemHandler *commands.RemoveItemFromCartHandler
}

func NewRabbitMQConsumer(channel *amqp091.Channel, queueName string, repo repositories.CartRepository) queue.QueueConsumer {
	addItemHandler := commands.NewAddItemToCartHandler(repo)
	removeItemHandler := commands.NewRemoveItemFromCartHandler(repo)
	return &RabbitMQConsumer{
		channel:           channel,
		queueName:         queueName,
		addItemHandler:    addItemHandler,
		removeItemHandler: removeItemHandler,
	}
}

func (c *RabbitMQConsumer) StartConsuming(ctx context.Context, handler func(message []byte) error) error {
	msgs, err := c.channel.Consume(
		c.queueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	go func() {
		for d := range msgs {
			if err := handler(d.Body); err != nil {
				log.Printf("Error handling message: %v", err)
			}
		}
	}()

	<-ctx.Done()
	return ctx.Err()
}

func (c *RabbitMQConsumer) HandleMessage(message []byte) error {
	var addItemCmd commands.AddItemToCartCommand
	var removeItemCmd commands.RemoveItemFromCartCommand

	if err := json.Unmarshal(message, &addItemCmd); err == nil {
		return c.addItemHandler.Handle(context.Background(), addItemCmd)
	} else if err := json.Unmarshal(message, &removeItemCmd); err == nil {
		return c.removeItemHandler.Handle(context.Background(), removeItemCmd)
	} else {
		return err
	}
}
