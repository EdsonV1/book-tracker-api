package handlers

import (
	"book-tracker-api/database"
	"book-tracker-api/models"
	"os"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestHandlers() *fiber.App {

	os.Setenv("JWT_SECRET", "test_secret_jwt_key")

	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to connect test database")
	}

	database.DB = db
	db.AutoMigrate(&models.Book{})
	db.AutoMigrate(&models.User{})

	app := fiber.New()

	app.Post("/books", CreateBook)
	app.Put("/books/:id", UpdateBook)
	app.Get("/books", GetBooks)
	app.Get("/books/:id", GetBook)
	app.Delete("/books/:id", DeleteBook)

	app.Post("/auth/register", Register)
	app.Post("/auth/login", Login)

	return app
}
