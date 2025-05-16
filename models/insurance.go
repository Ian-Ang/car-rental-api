package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Insurance struct {
	ID          uuid.UUID `gorm:"type:char(36);primary_key" json:"id"`
	Name        string    `gorm:"unique;not null" json:"name" example:"name"`
	Description string    `gorm:"not null" json:"description  " example:"description"`
	Price       float64   `gorm:"column:price; not null" json:"price" example:"50"`
	BaseModel
}

func (u *Insurance) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}
