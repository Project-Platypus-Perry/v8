package model

import (
	"time"

	"github.com/project-platypus-perry/v8/internal/constants"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID             string             `json:"ID" gorm:"primaryKey"`
	OrganizationID string             `json:"OrganizationID" gorm:"not null"`
	Organization   Organization       `json:"Organization" gorm:"foreignKey:OrganizationID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Role           constants.UserRole `json:"Role" gorm:"type:user_role"`
	Name           string             `json:"Name" gorm:"not null" validate:"required,min=2,max=50"`
	Email          string             `json:"Email" gorm:"not null;unique" validate:"required,email"`
	Password       string             `json:"Password" gorm:"not null" validate:"required,min=8"`
	Phone          string             `json:"Phone" gorm:"not null;unique" validate:"required,e164"`
	DateOfBirth    time.Time          `json:"DateOfBirth" gorm:"null" validate:"omitempty"`
}
