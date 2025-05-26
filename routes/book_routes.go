package routes

import (
	"book-tracker-api/handlers"
	"os"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func RegisterBookRoutes(app *fiber.App) {
	api := app.Group("/books")

	api.Get("/", handlers.GetBooks)
	api.Get("/:id", handlers.GetBook)

	api.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}))

	api.Post("/", handlers.CreateBook)
	api.Put("/:id", handlers.UpdateBook)
	api.Delete("/:id", handlers.DeleteBook)
}
