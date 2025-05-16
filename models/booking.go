package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Booking struct {
	ID        uuid.UUID `gorm:"type:char(36);primary_key" json:"id"`
	UserID    uuid.UUID `gorm:"column:user_id;not null" json:"user_id"`
	User      User
	CarID     uuid.UUID `gorm:"column:car_id;not null" json:"car_id"`
	Car       Car
	StartDate string `gorm:"not null" json:"start_date" example:"start_date"`
	EndDate   string `gorm:"not null" json:"end_date" example:"end_date"`
	Status    string `gorm:"default:'pending'" json:"status" example:"pending"` // "pending", "confirmed", "cancelled"
	BaseModel
}

func (u *Booking) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}
