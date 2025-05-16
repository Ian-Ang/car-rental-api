package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Role struct {
	ID    uuid.UUID `gorm:"type:char(36);primary_key" json:"id"`
	Name  string    `gorm:"uniqueIndex; not null" json:"name" example:"admin, customer"`
	Users []User
	BaseModel
}

func (r *Role) BeforeCreate(tx *gorm.DB) (err error) {
	r.ID = uuid.New()
	return
}
