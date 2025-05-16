package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CarCategory struct {
	ID   uuid.UUID `gorm:"type:char(36);primary_key" json:"id"`
	Name string    `gorm:"unique" json:"name" example:"name"` // e.g., "SUV", "Sedan", "Hatchback"
	Cars []Car     `gorm:"foreignKey:CategoryID" json:"cars"`
	BaseModel
}

func (u *CarCategory) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}
