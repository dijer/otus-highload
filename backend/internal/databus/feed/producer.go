package databus_feed

import (
	"encoding/json"

	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	exchange   = "posts"
	routingKey = "post.created"
)

func ProducePostCreated(ch *amqp.Channel, event PostCreatedEvent) error {
	body, err := json.Marshal(event)
	if err != nil {
		return err
	}

	err = ch.Publish(
		exchange,
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		return err
	}

	println("publish event", event.PostID)
	return nil
}
