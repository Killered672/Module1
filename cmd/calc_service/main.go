package main

import (
	"calc_service/internal/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/api/v1/calculate", handlers.CalculateExpression)

	println("Server is running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
