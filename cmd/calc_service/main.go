package main

import (
	"log"
	"net/http"

	"calc_service/internal/handlers"
)

func main() {
	http.HandleFunc("/api/v1/calculate", handlers.CalculateHandler)

	log.Println("Starting calculator service on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
