package main

import (
	"log"
	"net/http"

	"calc_service/internal/agent"
	"calc_service/internal/handlers"
	"calc_service/internal/orchestrator"
)

func main() {
	orchestratorInstance := orchestrator.NewOrchestrator()
	agent := agent.NewAgent(orchestratorInstance, 4)
	agent.Start()

	http.HandleFunc("/api/v1/calculate", handlers.CalculateHandler)

	log.Println("Starting calculator service on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
