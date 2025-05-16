package dto

import "github.com/google/uuid"

type CreateReviewInput struct {
	CarID   uuid.UUID `json:"car_id" binding:"required"`
	Rating  int       `json:"rating" binding:"required,min=1,max=5"`
	Comment string    `json:"comment"`
}

type ReviewOutput struct {
	ID      uuid.UUID `json:"id"`
	Car     string    `json:"car"`
	User    string    `json:"user"`
	Rating  int       `json:"rating"`
	Comment string    `json:"comment"`
}
