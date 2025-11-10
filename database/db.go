package database

import (
	"log"
	"os"

	
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectToDatabase() {
	dsn := os.Getenv("DB_CONNECTION_STRING")
	DB, err = gorm.Open(postgres.Open(dsn))

	if err != nil {
		log.Fatal("Error connecting to database")
	}
}
