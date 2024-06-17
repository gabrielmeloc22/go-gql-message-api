package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"

	"github.com/gabrielmeloc22/go-message-api/internal/config"
	"github.com/gabrielmeloc22/go-message-api/internal/context"
	"github.com/gabrielmeloc22/go-message-api/internal/handlers"
)

const DEFAULT_PORT = "8080"

var (
	Db *gorm.DB
)

func init() {
	godotenv.Load(".env")

	config.ConnectDb()
	config.Migrate()
}

func main() {

	port := DEFAULT_PORT
	if val, ok := os.LookupEnv("PORT"); ok {
		port = val
	}

	r := gin.Default()

	r.Use(context.GinContextToContextMiddleware())
	r.Use(context.Auth())

	r.POST("/query", handlers.GraphqlHandler())
	r.GET("/", handlers.PlaygroundHandler())

	r.Run(":" + port)
}
