package models

import (
	"time"
)

type Post struct {
	ID            uint   `gorm:"primaryKey"`
	Title         string `gorm:"not null"`
	Content       string `gorm:"type:text;not null"`
	AuthorID      uint   `gorm:"not null"`
	Author        User   `gorm:"foreignKey:AuthorID;constraint:OnDelete:CASCADE;"`
	AllowComments bool   `gorm:"default:true"`
	CreatedAt     time.Time
}
