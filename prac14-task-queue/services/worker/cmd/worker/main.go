package main

import (
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"example.com/prac14-task-queue/internal/amqp"
	"example.com/prac14-task-queue/internal/jobs"
	"example.com/prac14-task-queue/internal/store"
	amqplib "github.com/rabbitmq/amqp091-go"
)

func main() {
	rabbitURL := os.Getenv("RABBIT_URL")
	if rabbitURL == "" {
		rabbitURL = "amqp://guest:guest@localhost:5672/"
	}

	conn, err := amqplib.Dial(rabbitURL)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open channel: %v", err)
	}
	defer ch.Close()

	if err := amqp.DeclareQueues(ch); err != nil {
		log.Fatalf("Failed to declare queues: %v", err)
	}

	// Prefetch = 1
	if err := ch.Qos(1, 0, false); err != nil {
		log.Fatalf("Failed to set QoS: %v", err)
	}

	msgs, err := ch.Consume(amqp.MainQueue, "", false, false, false, false, nil)
	if err != nil {
		log.Fatalf("Failed to register consumer: %v", err)
	}

	publisher, err := amqp.NewPublisher(conn)
	if err != nil {
		log.Fatalf("Failed to create publisher: %v", err)
	}
	defer publisher.Close()

	processed := store.NewProcessedStore()

	log.Println("Worker started, waiting for jobs...")

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		for d := range msgs {
			var job jobs.TaskJob
			if err := json.Unmarshal(d.Body, &job); err != nil {
				log.Printf("Failed to unmarshal: %v", err)
				d.Nack(false, false)
				continue
			}

			log.Printf("Received job: task_id=%s, attempt=%d, message_id=%s",
				job.TaskID, job.Attempt, job.MessageID)

			// Идемпотентность: проверка, не обработано ли уже
			if processed.Exists(job.MessageID) {
				log.Printf("Duplicate message, ack without processing: %s", job.MessageID)
				d.Ack(false)
				continue
			}

			// Имитация обработки задачи
			err := processTask(job)

			if err != nil {
				log.Printf("Processing error: %v", err)
				job.Attempt++

				if job.Attempt <= amqp.MaxAttempts {
					// Повторная попытка
					if pubErr := publisher.PublishJob(amqp.MainQueue, job); pubErr != nil {
						log.Printf("Failed to publish retry: %v", pubErr)
					}
					d.Ack(false)
					continue
				}

				// Превышено число попыток — отправляем в DLQ
				log.Printf("Max attempts exceeded, sending to DLQ: %s", job.MessageID)
				if pubErr := publisher.PublishJob(amqp.DLQ, job); pubErr != nil {
					log.Printf("Failed to publish to DLQ: %v", pubErr)
				}
				d.Ack(false)
				continue
			}

			// Успешная обработка
			processed.MarkDone(job.MessageID)
			d.Ack(false)
			log.Printf("Job completed successfully: task_id=%s", job.TaskID)
		}
	}()

	<-done
	log.Println("Worker stopped")
}

func processTask(job jobs.TaskJob) error {
	// Имитация долгой обработки
	time.Sleep(2 * time.Second)

	// Искусственная ошибка для тестирования
	if job.TaskID == "t_fail" {
		return os.ErrDeadlineExceeded
	}
	return nil
}
