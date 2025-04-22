package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID          string    `json:"ID" gorm:"primaryKey"`
	Name        string    `json:"Name" gorm:"not null" validate:"required,min=2,max=50"`
	Email       string    `json:"Email" gorm:"not null;unique" validate:"required,email"`
	Password    string    `json:"Password" gorm:"not null" validate:"required,min=8"`
	Phone       string    `json:"Phone" gorm:"not null;unique" validate:"required,e164"`
	DateOfBirth time.Time `json:"DateOfBirth" gorm:"null" validate:"omitempty"`
}
