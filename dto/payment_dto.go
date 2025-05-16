package dto

import "github.com/google/uuid"

type CreatePaymentInput struct {
	BookingID uuid.UUID `json:"booking_id" binding:"required"`
	Amount    float64   `json:"amount" binding:"required"`
	Method    string    `json:"method" binding:"required"`
}

type PaymentOutput struct {
	ID        uuid.UUID `json:"id"`
	BookingID uuid.UUID `json:"booking_id"`
	Amount    float64   `json:"amount"`
	Method    string    `json:"method"`
	Status    string    `json:"status"`
}
