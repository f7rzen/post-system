package main

import (
	"github.com/joho/godotenv"
	"log"
	"post-system/pkg/db"
)

func init() {
	// Загрузка переменных окружения
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}
	db.ConnectToDb()
	db.MigrationDatabase()
}

func main() {

}
