package amqp

import (
	"context"
	"encoding/json"
	"log"

	"example.com/prac13-rabbit/internal/events"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Publisher struct {
	channel   *amqp.Channel
	queueName string
}

func NewPublisher(conn *amqp.Connection, queueName string) (*Publisher, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	_, err = ch.QueueDeclare(
		queueName,
		true, false, false, false, nil,
	)
	if err != nil {
		return nil, err
	}

	return &Publisher{channel: ch, queueName: queueName}, nil
}

func (p *Publisher) PublishTaskCreated(taskID, requestID string) error {
	event := events.NewTaskCreatedEvent(taskID, requestID)
	body, err := json.Marshal(event)
	if err != nil {
		return err
	}

	err = p.channel.PublishWithContext(
		context.Background(),
		"",
		p.queueName,
		false,
		false,
		amqp.Publishing{
			ContentType:  "application/json",
			DeliveryMode: amqp.Persistent,
			Body:         body,
		},
	)
	if err != nil {
		log.Printf("Failed to publish: %v", err)
	}
	return err
}

func (p *Publisher) Close() error {
	return p.channel.Close()
}
