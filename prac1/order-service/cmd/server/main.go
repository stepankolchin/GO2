package main

import (
	"log"
	"net/http"
	"strings"
	"fmt"

	"example.com/order-service/internal/order"
)

func main() {
	repo := order.NewRepo()
	client := order.NewUserServiceClient("http://localhost:8081")
	handler := order.NewHandler(repo, client)

	mux := http.NewServeMux()

	mux.HandleFunc("/orders/", func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/full") {
			fmt.Println("full handle")
			handler.GetOrderWithUser(w, r)
			return
		}
		if strings.HasPrefix(r.URL.Path, "/orders/by-user") {
			fmt.Println("by-user handle")
			handler.GetAllUserOrders(w, r)
			return
		}

		handler.GetOrderByID(w, r)
	})

	addr := ":8082"
	log.Println("order-service started on", addr)

	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}

