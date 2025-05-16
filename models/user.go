package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	CreatedAt time.Time      `gorm:"->:false;column:created_at" json:"-"`
	UpdatedAt time.Time      `gorm:"->:false;column:updated_at" json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"->:false;column:deleted_at" json:"-"`
}

type User struct {
	ID       uuid.UUID `gorm:"type:char(36);primary_key" json:"id"`
	Name     string    `gorm:"not null" json:"name" example:"Rian"`
	Email    string    `gorm:"unique; not null" json:"email" example:"Rian@gmail.com"`
	Password string    `gorm:"not null" json:"password" example:"P@ssW0rd"`
	RoleID   uuid.UUID `gorm:"column:role_id;not null" json:"role_id"`
	Role     Role      `gorm:"foreignKey:RoleID;references:ID" json:"role"`
	Cars     []Car     `gorm:"foreignKey:OwnerID" json:"cars"`
	Bookings []Booking
	Reviews  []Review
	BaseModel
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}
