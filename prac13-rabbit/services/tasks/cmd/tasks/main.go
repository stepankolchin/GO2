package main

import (
	"log"
	"net/http"
	"os"

	"example.com/prac13-rabbit/internal/amqp"
	"example.com/prac13-rabbit/internal/httpapi"
	"example.com/prac13-rabbit/internal/task"
	amqplib "github.com/rabbitmq/amqp091-go"
)

func main() {
	rabbitURL := os.Getenv("RABBIT_URL")
	if rabbitURL == "" {
		rabbitURL = "amqp://guest:guest@localhost:5672/"
	}
	queueName := os.Getenv("QUEUE_NAME")
	if queueName == "" {
		queueName = "task_events"
	}

	conn, err := amqplib.Dial(rabbitURL)
	if err != nil {
		log.Fatal("Failed to connect to RabbitMQ:", err)
	}
	defer conn.Close()

	publisher, err := amqp.NewPublisher(conn, queueName)
	if err != nil {
		log.Fatal("Failed to create publisher:", err)
	}
	defer publisher.Close()

	repo := task.NewRepo()
	handler := httpapi.NewHandler(repo, publisher)

	mux := http.NewServeMux()
	mux.HandleFunc("/v1/tasks", handler.CreateTask)

	log.Println("Tasks service started on :8082")
	log.Fatal(http.ListenAndServe(":8082", mux))
}
