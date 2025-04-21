package model

import (
	"time"

	"gorm.io/gorm"
)

type Batch struct {
	gorm.Model
	ID             string       `json:"id" gorm:"primaryKey"`
	OrganizationID string       `json:"organization_id" gorm:"not null"`
	Organization   Organization `json:"organization" gorm:"foreignKey:OrganizationID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Name           string       `json:"name" gorm:"not null"`
	Description    string       `json:"description" gorm:"null"`
	CreatedAt      time.Time    `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt      time.Time    `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt      time.Time    `json:"deleted_at" gorm:"index"`
}
