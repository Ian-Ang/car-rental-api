package dto

import "github.com/google/uuid"

type CreateMaintenanceInput struct {
	CarID   uuid.UUID `json:"car_id" binding:"required"`
	Details string    `json:"details" binding:"required"`
	Date    string    `json:"date" binding:"required,datetime=2006-01-02"`
}

type MaintenanceOutput struct {
	ID      uuid.UUID `json:"id"`
	Car     string    `json:"car"`
	Details string    `json:"details"`
	Date    string    `json:"date"`
}
