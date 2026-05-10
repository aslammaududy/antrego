package config

import (
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	db, err := gorm.Open(sqlite.Open(os.Getenv("GOOSE_DBSTRING")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB = db
}
