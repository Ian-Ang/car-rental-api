package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Car struct {
	ID                uuid.UUID   `gorm:"type:char(36);primary_key" json:"id"`
	Brand             string      `gorm:"not null" json:"brand" example:"brand"`
	Model             string      `gorm:"not null" json:"model" example:"model"`
	Year              string      `gorm:"not null" json:"year" example:"2000"`
	LicensePlate      string      `gorm:"not null" json:"licenseplate" example:"B2345AJA"`
	PricePerDay       float64     `gorm:"column:price_per_day; not null" json:"price_per_day" example:"50000000"`
	Description       string      `gorm:"not null" json:"description"`
	CategoryID        uuid.UUID   `gorm:"column:category_id;not null" json:"category_id"`
	Category          CarCategory `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"category"`
	OwnerID           uuid.UUID   `gorm:"column:owner_id;not null" json:"owner_id"`
	Owner             User
	PickupLocationID  uuid.UUID `gorm:"column:location_id;not null" json:"location_id"`
	PickupLocation    Location
	DropoffLocationID uuid.UUID `gorm:"column:dropoff_location_id;not null" json:"dropoff_location_id"`
	DropoffLocation   Location
	Status            string `gorm:"default:'approved'" json:"status" example:"approved"` // "pending_approval", "approved", "rejected", "rented", "maintenance"
	BaseModel
}

func (u *Car) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}
