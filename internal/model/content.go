package model

import (
	"github.com/gagan-gaurav/base/internal/constants"
	"gorm.io/gorm"
)

type Content struct {
	gorm.Model
	ID          string                `json:"ID" gorm:"primaryKey"`
	ChapterID   string                `json:"ChapterID" gorm:"not null"`
	Chapter     Chapter               `json:"Chapter" gorm:"foreignKey:ChapterID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Type        constants.ContentType `json:"Type" gorm:"type:content_type;default:'notes'"`
	Name        string                `json:"Name" gorm:"not null"`
	Description string                `json:"Description" gorm:"null"`
	Language    constants.Language    `json:"Language" gorm:"type:language;not null;default:'en'"`
	Visibility  constants.Visibility  `json:"Visibility" gorm:"type:visibility;not null;default:'public'"`
}
