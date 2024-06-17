package config

import (
	"log"
	"os"

	"github.com/gabrielmeloc22/go-message-api/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	Db *gorm.DB
)

func ConnectDb() {
	dsn, ok := os.LookupEnv("DATABASE_CONFIG")

	if !ok {
		log.Fatal("connection string not found")
	}

	db, err := gorm.Open(postgres.Open(dsn))

	if err != nil {
		log.Fatal("fail to connect database")
	}
	Db = db
}

func Migrate() {
	Db.AutoMigrate(model.Message{}, model.User{}, model.Chat{})
}
