package database

import (
	"book-tracker-api/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(sqlite.Open("gom.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to DB")
	}

	db.AutoMigrate(&models.Book{})
	db.AutoMigrate(&models.User{})

	DB = db
}
