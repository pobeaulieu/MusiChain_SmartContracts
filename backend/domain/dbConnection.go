package domain

import (
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	err := godotenv.Load("environment.env")
	if err != nil {
		panic("could not load env variables")
	}
	dbURL := os.Getenv("DB_URL")
	connection, err := gorm.Open(sqlite.Open(dbURL), &gorm.Config{})

	if err != nil {
		panic("could not connect to the database")
	}

	DB = connection

	err = connection.AutoMigrate(&User{})
	if err != nil {
		panic("could not migrate the database")
	}
}
