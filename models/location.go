package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Location struct {
	ID       uuid.UUID `gorm:"type:char(36);primary_key" json:"id"`
	City     string    `gorm:"not null" json:"city" example:"city"`
	Address  string    `gorm:"not null" json:"adddress" example:"adddress"`
	Cars     []Car     `gorm:"foreignKey:PickupLocationID"`
	Dropoffs []Car     `gorm:"foreignKey:DropoffLocationID"`
	BaseModel
}

func (u *Location) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}
