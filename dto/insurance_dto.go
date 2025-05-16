package dto

import "github.com/google/uuid"

type CreateInsuranceInput struct {
	Name   string  `json:"name" binding:"required"`
	Amount float64 `json:"amount" binding:"required"`
}

type InsuranceOutput struct {
	ID     uuid.UUID `json:"id"`
	Name   string    `json:"name"`
	Amount float64   `json:"amount"`
}
