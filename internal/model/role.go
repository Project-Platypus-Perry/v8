package model

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	UserID         string       `json:"user_id" gorm:"primaryKey;type:uuid;not null"`
	OrganizationID string       `json:"organization_id" gorm:"primaryKey;type:uuid;not null"`
	Role           string       `json:"role" gorm:"not null"`
	User           User         `json:"user" gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Organization   Organization `json:"organization" gorm:"foreignKey:OrganizationID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedAt      time.Time    `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt      time.Time    `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt      time.Time    `json:"deleted_at" gorm:"index"`
}
