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

type UsersClassrooms struct {
	gorm.Model
	// Composite key on UserID and ClassroomID
	UserID         string       `json:"UserID" gorm:"primaryKey"`
	User           User         `json:"User" gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	ClassroomID    string       `json:"ClassroomID" gorm:"primaryKey"`
	Classroom      Classroom    `json:"Classroom" gorm:"foreignKey:ClassroomID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	OrganizationID string       `json:"OrganizationID" gorm:"not null"`
	Organization   Organization `json:"Organization" gorm:"foreignKey:OrganizationID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	BatchID        string       `json:"BatchID" gorm:"not null"`
	Batch          Batch        `json:"Batch" gorm:"foreignKey:BatchID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
