package model

// import (
// 	"time"

// 	"github.com/gagan-gaurav/base/internal/constants"
// 	"gorm.io/gorm"
// )

// type Content struct {
// 	gorm.Model
// 	ID          string                `json:"id" gorm:"primaryKey"`
// 	ChapterID   string                `json:"chapter_id" gorm:"not null"`
// 	Chapter     Chapter               `json:"chapter" gorm:"foreignKey:ChapterID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
// 	Type        constants.ContentType `json:"type" gorm:"type:content_type;default:'notes'"`
// 	Name        string                `json:"name" gorm:"not null"`
// 	Description string                `json:"description" gorm:"null"`
// 	Language    constants.Language    `json:"language" gorm:"type:language;not null;default:'en'"`
// 	Visibility  constants.Visibility  `json:"visibility" gorm:"type:visibility;not null;default:'public'"`
// 	CreatedAt   time.Time             `json:"created_at" gorm:"autoCreateTime"`
// 	UpdatedAt   time.Time             `json:"updated_at" gorm:"autoUpdateTime"`
// 	DeletedAt   gorm.DeletedAt        `json:"deleted_at" gorm:"index"`
// }
