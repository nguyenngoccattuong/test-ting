package Initializers

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error

	// Sử dụng SQLite thay vì PostgreSQL
	DB, err = gorm.Open(sqlite.Open("kol.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}
}
