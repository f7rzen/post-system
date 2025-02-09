package graph

import (
	"gorm.io/gorm"
)

// Resolver - основная структура, которая хранит зависимости
type Resolver struct {
	DB *gorm.DB // Здесь должно быть твое подключение к базе данных
}
