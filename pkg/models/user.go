package models

import (
	"time"
)

// User - таблица пользователей
type User struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"unique;not null"`
	CreatedAt time.Time
}
