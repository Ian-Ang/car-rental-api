package dto

import "github.com/google/uuid"

type CreateCategoryInput struct {
	Name string `json:"name" binding:"required"`
}

type CategoryOutput struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
