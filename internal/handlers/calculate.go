package handlers

import (
	"calc_service/internal/models"
	"calc_service/internal/services"
	"encoding/json"
	"net/http"
)

func CalculateExpression(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req models.CalculateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	result, err := services.Calculate(req.Expression)
	if err != nil {
		if err.Error() == "invalid expression" {
			http.Error(w, "Expression is not valid", http.StatusUnprocessableEntity)
			return
		}
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"result": result})
}
