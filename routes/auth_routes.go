package routes

import (
	"book-tracker-api/handlers"

	"github.com/gofiber/fiber/v2"
)

func RegisterAuth(app *fiber.App) {
	auth := app.Group("/auth")

	auth.Post("/register", handlers.Register)
	auth.Post("/login", handlers.Login)
}
