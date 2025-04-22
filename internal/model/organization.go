package model

import (
	"gorm.io/gorm"
)

type Organization struct {
	gorm.Model
	ID          string `json:"ID" gorm:"primaryKey"`
	Name        string `json:"Name" gorm:"not null"`
	Description string `json:"Description" gorm:"not null"`
}
