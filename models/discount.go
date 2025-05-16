package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Discount struct {
	ID         uuid.UUID `gorm:"type:char(36);primary_key" json:"id"`
	Name       string    `gorm:"unique;not null" json:"name" example:"name"`
	Percentage float64   `gorm:"column:percentage ; not null" json:"percentage " example:"50"`
	ExpiresAt  string    `gorm:"not null" json:"expires_at  " example:"ExpiresAt "`
	BaseModel
}

func (u *Discount) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}
