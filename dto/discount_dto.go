package dto

import "github.com/google/uuid"

type CreateDiscountInput struct {
	Code       string  `json:"code" binding:"required"`
	Percentage float64 `json:"percentage" binding:"required"`
}

type DiscountOutput struct {
	ID         uuid.UUID `json:"id"`
	Code       string    `json:"code"`
	Percentage float64   `json:"percentage"`
}
