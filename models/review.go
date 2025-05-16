package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Review struct {
	ID      uuid.UUID `gorm:"type:char(36);primary_key" json:"id"`
	UserID  uuid.UUID `gorm:"column:user_id;not null" json:"user_id"`
	User    User
	CarID   uuid.UUID `gorm:"column:car_id;not null" json:"car_id"`
	Car     Car
	Rating  int    `gorm:"column:rating" json:"rating" example:"rating"`
	Comment string `gorm:"not null" json:"comment" example:"comment"`
	BaseModel
}

func (u *Review) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}
