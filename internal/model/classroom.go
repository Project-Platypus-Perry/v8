package model

import (
	"time"

	"gorm.io/gorm"
)

type Classroom struct {
	gorm.Model
	ID             string         `json:"id" gorm:"primaryKey"`
	OrganizationID string         `json:"organization_id" gorm:"not null"`
	Organization   Organization   `json:"organization" gorm:"foreignKey:OrganizationID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	BatchID        string         `json:"batch_id" gorm:"not null"`
	Batch          Batch          `json:"batch" gorm:"foreignKey:BatchID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Name           string         `json:"name" gorm:"not null"`
	Description    string         `json:"description" gorm:"null"`
	Settings       string         `json:"settings" gorm:"type:jsonb"`
	CreatedAt      time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt      time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt      gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
