package handlers

import (
	"book-tracker-api/database"
	"book-tracker-api/models"
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func getJwtSecret() []byte {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		panic("JWT_SECRET environment variable not set")
	}
	return []byte(secret)
}

func Register(c *fiber.Ctx) error {
	var data models.User

	if err := c.BodyParser(&data); err != nil {
		fmt.Println("DB Error:", err.Error())
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	hashed, _ := bcrypt.GenerateFromPassword([]byte(data.Password), 14)
	data.Password = string(hashed)

	if err := database.DB.Create(&data).Error; err != nil {
		fmt.Println("DB Error:", err.Error())
		return c.Status(400).JSON(fiber.Map{"error": "Username already exists"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "User registered!"})
}

func Login(c *fiber.Ctx) error {
	var data models.User
	var user models.User

	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	database.DB.Where("username = ?", data.Username).First(&user)

	if user.ID == 0 {
		return c.Status(400).JSON(fiber.Map{"error": "User not found"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password)); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Wrong password"})
	}

	claims := jwt.MapClaims{
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString(getJwtSecret())

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not sign token"})
	}

	return c.JSON(fiber.Map{"token": signed})
}
