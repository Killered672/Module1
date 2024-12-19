package models

type CalculateRequest struct {
	Expression string `json:"expression" binding:"required"`
}
