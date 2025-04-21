package model

import (
	"time"

	"gorm.io/gorm"
)

type Chapter struct {
	gorm.Model
	ID          string         `json:"id" gorm:"primaryKey"`
	ClassroomID string         `json:"classroom_id" gorm:"not null"`
	Classroom   Classroom      `json:"classroom" gorm:"foreignKey:ClassroomID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Name        string         `json:"name" gorm:"not null"`
	Description string         `json:"description" gorm:"null"`
	CreatedAt   time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
