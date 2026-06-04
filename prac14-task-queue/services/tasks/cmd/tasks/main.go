package main

import (
	"log"
	"net/http"
	"os"

	"example.com/prac14-task-queue/internal/amqp"
	"example.com/prac14-task-queue/internal/httpapi"
	amqplib "github.com/rabbitmq/amqp091-go"
)

func main() {
	rabbitURL := os.Getenv("RABBIT_URL")
	if rabbitURL == "" {
		rabbitURL = "amqp://guest:guest@localhost:5672/"
	}

	conn, err := amqplib.Dial(rabbitURL)
	if err != nil {
		log.Fatal("Failed to connect to RabbitMQ:", err)
	}
	defer conn.Close()

	publisher, err := amqp.NewPublisher(conn)
	if err != nil {
		log.Fatal("Failed to create publisher:", err)
	}
	defer publisher.Close()

	handler := httpapi.NewHandler(publisher)

	mux := http.NewServeMux()
	mux.HandleFunc("/v1/jobs/process-task", handler.SubmitJob)

	log.Println("Tasks service started on :8082")
	log.Fatal(http.ListenAndServe(":8082", mux))
}
