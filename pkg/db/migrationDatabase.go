package db

import (
	"log"

	"post-system/pkg/models"
)

func MigrationDatabase() {
	if err := DB.AutoMigrate(&models.Comment{}); err != nil {
		log.Fatalf("ошибка миграции таблицы Comment: %v", err)
	}
	if err := DB.AutoMigrate(&models.Post{}); err != nil {

		log.Fatalf("ошибка миграции таблицы Post: %v", err)
	}
	if err := DB.AutoMigrate(&models.User{}); err != nil {

		log.Fatalf("ошибка миграции таблицы User: %v", err)
	}
}
