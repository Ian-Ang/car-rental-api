package dto

import "github.com/google/uuid"

type CreateBookingInput struct {
	CarID     uuid.UUID `json:"car_id" binding:"required"`
	StartDate string    `json:"start_date" binding:"required,datetime=2006-01-02"`
	EndDate   string    `json:"end_date" binding:"required,datetime=2006-01-02"`
}

type BookingOutput struct {
	ID        uuid.UUID `json:"id"`
	Car       string    `json:"car"`
	User      string    `json:"user"`
	StartDate string    `json:"start_date"`
	EndDate   string    `json:"end_date"`
	Status    string    `json:"status"`
}
