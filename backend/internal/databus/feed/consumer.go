package databus_feed

import (
	"context"
	"encoding/json"
	"fmt"

	infra_database "github.com/dijer/otus-highload/backend/internal/infra/database"
	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	queue = "feed.worker"
)

func ConsumePostCreated(ctx context.Context, ch *amqp.Channel, dbRouter *infra_database.DBRouter) error {
	q, err := ch.QueueDeclare(
		queue,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	if err := ch.QueueBind(q.Name, routingKey, exchange, false, nil); err != nil {
		return err
	}

	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		return err
	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case d := <-msgs:
				var e PostCreatedEvent
				if err := json.Unmarshal(d.Body, &e); err != nil {
					println("invalid message")
					continue
				}

				rows, err := dbRouter.Query(ctx, `
					SELECT user_id FROM follows WHERE friend_id=$1
				`, e.AuthorUserID)
				if err != nil {
					continue
				}

				var subs []int64
				for rows.Next() {
					var userID int64
					rows.Scan(&userID)
					subs = append(subs, userID)
				}
				rows.Close()

				for _, userID := range subs {
					body, _ := json.Marshal(e)
					routingKey := fmt.Sprintf("feed.deliver.%d", userID)
					ch.Publish("feeds", routingKey, false, false, amqp.Publishing{
						ContentType: "application/json",
						Body:        body,
					})
				}
			}
		}
	}()

	return nil
}
