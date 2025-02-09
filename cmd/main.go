package main

import (
	"log"
	"post-system/graph"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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
	// Создаём сервер на Gin
	r := gin.Default()

	// Создаём GraphQL сервер
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	// GraphQL Playground (UI для тестов)
	r.GET("/", func(c *gin.Context) {
		playground.Handler("GraphQL Playground", "/query").ServeHTTP(c.Writer, c.Request)
	})

	// GraphQL API
	r.POST("/query", func(c *gin.Context) {
		srv.ServeHTTP(c.Writer, c.Request)
	})

	// Запуск сервера
	log.Println("Server running on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
