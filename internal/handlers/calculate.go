package handlers

import (
	"encoding/json"
	"net/http"

	"calc_service/internal/evaluator"
	"calc_service/internal/models"
)

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

	result, err := evaluator.EvaluateExpression(req.Expression)
	if err != nil {
		if err.Error() == "invalid expression" {
			sendErrorResponse(w, http.StatusUnprocessableEntity, "Expression is not valid")
		} else {
			sendErrorResponse(w, http.StatusInternalServerError, "Internal server error")
		}
		return
	}

	response := models.CalculateResponse{Result: result}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func sendErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}
