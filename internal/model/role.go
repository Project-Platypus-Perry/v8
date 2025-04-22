package model

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	UserID         string       `json:"UserID" gorm:"primaryKey;type:uuid;not null"`
	OrganizationID string       `json:"OrganizationID" gorm:"primaryKey;type:uuid;not null"`
	Role           string       `json:"Role" gorm:"not null"`
	User           User         `json:"User" gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Organization   Organization `json:"Organization" gorm:"foreignKey:OrganizationID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
