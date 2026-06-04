package amqp

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	MainQueue   = "task_jobs"
	DLQ         = "task_jobs_dlq"
	MaxAttempts = 3
)

func DeclareQueues(ch *amqp.Channel) error {
	// Сначала объявляем DLQ
	_, err := ch.QueueDeclare(
		DLQ,
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return err
	}

	// Основная очередь с DLQ
	args := amqp.Table{
		"x-dead-letter-exchange":    "",
		"x-dead-letter-routing-key": DLQ,
	}
	_, err = ch.QueueDeclare(
		MainQueue,
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		args,
	)
	return err
}
