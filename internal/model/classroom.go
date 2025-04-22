package model

import (
	"gorm.io/gorm"
)

type Classroom struct {
	gorm.Model
	ID             string       `json:"ID" gorm:"primaryKey"`
	OrganizationID string       `json:"OrganizationID" gorm:"not null"`
	Organization   Organization `json:"Organization" gorm:"foreignKey:OrganizationID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	BatchID        string       `json:"BatchID" gorm:"not null"`
	Batch          Batch        `json:"Batch" gorm:"foreignKey:BatchID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Name           string       `json:"Name" gorm:"not null"`
	Description    string       `json:"Description" gorm:"null"`
	Settings       string       `json:"Settings" gorm:"type:jsonb"`
}
