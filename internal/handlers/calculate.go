package handlers

import (
	"encoding/json"
	"net/http"

	"calc_service/internal/models"
	"calc_service/internal/orchestrator"
)

var orchestratorInstance = orchestrator.NewOrchestrator()

func CalculateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var req models.CalculateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		sendErrorResponse(w, http.StatusBadRequest, "Invalid JSON format")
		return
	}

	id, err := orchestratorInstance.AddExpression(req.Expression)
	if err != nil {
		sendErrorResponse(w, http.StatusUnprocessableEntity, "Invalid expression")
		return
	}

	response := models.CalculateResponse{ID: id}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func sendErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}
