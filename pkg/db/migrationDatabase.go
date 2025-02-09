package db

import (
	"log"

	"post-system/pkg/models"
)

func MigrationDatabase() {
	if err := DB.AutoMigrate(&models.Product{}); err != nil {
		log.Fatalf("ошибка миграции таблицы Product: %v", err)
	}
	if err := DB.AutoMigrate(&models.User{}); err != nil {

		log.Fatalf("ошибка миграции таблицы User: %v", err)
	}
}
