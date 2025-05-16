package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Payment struct {
	ID        uuid.UUID `gorm:"type:char(36);primary_key" json:"id"`
	BookingID uuid.UUID `gorm:"column:booking_id;not null" json:"booking_id"`
	Booking   Booking
	Amount    float64 `gorm:"column:amount; not null" json:"amount" example:"50000000"`
	Method    string  `gorm:"not null" json:"method" example:"method"`
	Status    string  `gorm:"default:'paid'" json:"status" example:"paid"` // paid, failed, refunded
	BaseModel
}

func (u *Payment) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}
