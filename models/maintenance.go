package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Maintenance struct {
	ID          uuid.UUID `gorm:"type:char(36);primary_key" json:"id"`
	CarID       uuid.UUID `gorm:"column:car_id;not null" json:"car_id"`
	Car         Car
	Description string `gorm:"not null" json:"description" example:"description"`
	Date        string `gorm:"not null" json:"date" example:"date"`
	BaseModel
}

func (u *Maintenance) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}
