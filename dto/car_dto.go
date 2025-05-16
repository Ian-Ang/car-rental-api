package dto

import "github.com/google/uuid"

type CreateCarInput struct {
	Model        string    `json:"model" binding:"required" example:"model"`
	Year         string    `json:"year" binding:"required" example:"2000"`
	LicensePlate string    `json:"licenseplate" binding:"required" example:"B2345AJA"`
	Description  string    `json:"description" binding:"required"`
	PricePerDay  float64   `json:"price_per_day" binding:"required"`
	CategoryID   uuid.UUID `json:"category_id" binding:"required"`
	LocationID   uuid.UUID `json:"location_id" binding:"required"`
	OwnerID      uuid.UUID `json:"owner_id" binding:"required"`
	Status       string    `json:"status" binding:"required" example:"approved"`
}

type CarOutput struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	PricePerDay float64   `json:"price_per_day"`
	Category    string    `json:"category"`
	Location    string    `json:"location"`
}
