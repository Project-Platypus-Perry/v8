package model

import (
	"time"

	"gorm.io/gorm"
)

type Batch struct {
	gorm.Model
	ID             string       `json:"ID" gorm:"primaryKey"`
	OrganizationID string       `json:"OrganizationID" gorm:"not null"`
	Organization   Organization `json:"Organization" gorm:"foreignKey:OrganizationID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Name           string       `json:"Name" gorm:"not null" validate:"required"`
	Description    string       `json:"Description" gorm:"null" validate:"omitempty"`
}

type UsersBatches struct {
	// Composite key on UserID and BatchID
	UserID         string         `json:"UserID" gorm:"primaryKey"`
	User           User           `json:"User" gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	BatchID        string         `json:"BatchID" gorm:"primaryKey"`
	Batch          Batch          `json:"Batch" gorm:"foreignKey:BatchID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	OrganizationID string         `json:"OrganizationID" gorm:"not null"`
	Organization   Organization   `json:"Organization" gorm:"foreignKey:OrganizationID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedAt      time.Time      `json:"CreatedAt"`
	UpdatedAt      time.Time      `json:"UpdatedAt"`
	DeletedAt      gorm.DeletedAt `json:"DeletedAt" gorm:"index"`
}

type BatchResponseModel struct {
	ID             string `json:"ID" gorm:"primaryKey"`
	OrganizationID string `json:"OrganizationID" gorm:"not null"`
	Name           string `json:"Name" gorm:"not null"`
	Description    string `json:"Description" gorm:"null"`
}

// TableName specifies the table name for BatchResponseModel
func (BatchResponseModel) TableName() string {
	return "batches"
}

type AssociateUserToBatchRequest struct {
	BatchID string   `json:"BatchID" validate:"required"`
	UserIDs []string `json:"UserIDs" validate:"required"`
}
