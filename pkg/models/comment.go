package models

import (
	"time"
)

// Comment представляет комментарий к посту
type Comment struct {
	ID        uint      `gorm:"primaryKey"`
	PostID    uint      `gorm:"not null"`
	Post      Post      `gorm:"foreignKey:PostID;constraint:OnDelete:CASCADE;"`
	AuthorID  uint      `gorm:"not null"`
	Author    User      `gorm:"foreignKey:AuthorID;constraint:OnDelete:CASCADE;"`
	ParentID  *uint     `gorm:"index"`
	Parent    *Comment  `gorm:"foreignKey:ParentID;constraint:OnDelete:CASCADE;"`
	Content   string    `gorm:"type:text;not null"`
	CreatedAt time.Time
  }