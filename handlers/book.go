package handlers

import (
	"book-tracker-api/database"
	"book-tracker-api/models"

	"github.com/gofiber/fiber/v2"
)

func GetBooks(c *fiber.Ctx) error {
	readParam := c.Query("read")
	var books []models.Book

	if readParam == "true" {
		database.DB.Where("read = ?", true).Find(&books)
	} else if readParam == "false" {
		database.DB.Where("read = ?", false).Find(&books)
	} else {
		database.DB.Find(&books)
	}

	return c.JSON(books)
}

func GetBook(c *fiber.Ctx) error {
	id := c.Params("id")
	var book models.Book
	result := database.DB.First(&book, id)

	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Book not found"})
	}

	return c.JSON(book)
}

func CreateBook(c *fiber.Ctx) error {
	book := new(models.Book)

	if err := c.BodyParser(book); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	database.DB.Create(&book)

	return c.Status(201).JSON(book)
}

func UpdateBook(c *fiber.Ctx) error {
	id := c.Params("id")
	var book models.Book

	if err := database.DB.First(&book, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Book not found"})
	}

	var input models.Book
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	book.Title = input.Title
	book.Author = input.Author
	book.Read = input.Read

	if err := database.DB.Save(&book).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update book"})
	}

	return c.JSON(book)
}

func DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")
	result := database.DB.Delete(&models.Book{}, id)

	if result.RowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "Book not found"})
	}

	return c.SendStatus(204)
}
