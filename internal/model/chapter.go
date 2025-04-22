package model

import (
	"gorm.io/gorm"
)

type Chapter struct {
	gorm.Model
	ID          string    `json:"ID" gorm:"primaryKey"`
	ClassroomID string    `json:"ClassroomID" gorm:"not null"`
	Classroom   Classroom `json:"Classroom" gorm:"foreignKey:ClassroomID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Name        string    `json:"Name" gorm:"not null"`
	Description string    `json:"Description" gorm:"null"`
}
