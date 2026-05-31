package main

import (
	"log"
	"net/http"

	"example.com/user-service/internal/user"
)

func main() {
	repo := user.NewRepo()
	handler := user.NewHandler(repo)

	mux := http.NewServeMux()
	mux.HandleFunc("/users/", handler.GetUserByID)
	mux.HandleFunc("/users", handler.GetAllUsers)


	addr := ":8081"
	log.Println("user-service started on", addr)

	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}

