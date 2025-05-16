package dto

import "github.com/google/uuid"

type RoleInput struct {
	Name string `json:"name" binding:"required"`
}

type RoleOutput struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
