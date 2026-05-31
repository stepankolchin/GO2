package main

import (
	"context"
	"log"
	"time"

	"example.com/prac2-grpc/gen/studentpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient(
		"localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := studentpb.NewStudentServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Вызов метода Ping
	pingResp, err := client.Ping(ctx, &studentpb.PingRequest{
		Message: "hello grpc",
	})
	if err != nil {
		log.Fatal("Ping error:", err)
	}
	log.Println("Ping response:", pingResp.GetMessage())

	// Вызов метода GetStudentByID
	studentResp, err := client.GetStudentByID(ctx, &studentpb.GetStudentRequest{
		Id: 1,
	})
	if err != nil {
		log.Fatal("GetStudentByID error:", err)
	}

	st := studentResp.GetStudent()
	log.Printf("Student: id=%d, full_name=%s, group=%s, email=%s\n",
		st.GetId(),
		st.GetFullName(),
		st.GetGroup(),
		st.GetEmail(),
	)
}

