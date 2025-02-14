package main

import (
	"log"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"post-system/graph"
	"post-system/pkg/db"
)

func init() {
	// Загружаем переменные окружения
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}
	db.ConnectToDb()
	db.MigrationDatabase()
}

func main() {
	r := gin.Default()

	// Добавляем GraphQL эндпоинт

	// // GraphQL Playground для тестирования API
	// r.GET("/", gin.WrapH(playground.Handler("GraphQL", "/query")))

	resolver := &graph.Resolver{
		DB:            db.DB,
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))
	r.POST("/query", gin.WrapH(srv))

	log.Println("Сервер запущен на http://localhost:8080")
	r.Run(":8080")
}
