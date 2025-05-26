package main

import (
	"book-tracker-api/database"
	"book-tracker-api/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	app := fiber.New()
	database.Connect()

	routes.RegisterBookRoutes(app)
	routes.RegisterAuth(app)

	app.Listen(":3000")
}
