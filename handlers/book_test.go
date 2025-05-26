package handlers

import (
	"book-tracker-api/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetBooks(t *testing.T) {
	app := setupTestHandlers()
	book := models.Book{
		Title:  "Test Driven Development",
		Author: "Kent Beck",
		Read:   true,
	}
	payload, _ := json.Marshal(book)

	req := httptest.NewRequest(http.MethodPost, "/books", bytes.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
	defer resp.Body.Close()

	req = httptest.NewRequest(http.MethodGet, "/books", bytes.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")

	var books models.Book
	err = json.NewDecoder(resp.Body).Decode(&books)
	assert.NoError(t, err)

	assert.NotEmpty(t, books)
}

func TestCreateBook(t *testing.T) {
	app := setupTestHandlers()

	book := models.Book{
		Title:  "Test Driven Development",
		Author: "Kent Beck",
		Read:   true,
	}
	payload, _ := json.Marshal(book)

	req := httptest.NewRequest(http.MethodPost, "/books", bytes.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
	defer resp.Body.Close()

	var created models.Book
	err = json.NewDecoder(resp.Body).Decode(&created)
	assert.NoError(t, err)

	assert.Equal(t, book.Title, created.Title)
	assert.Equal(t, book.Author, created.Author)
	assert.Equal(t, book.Read, created.Read)
	assert.NotZero(t, created.ID)
}

func TestUpdateBook(t *testing.T) {
	app := setupTestHandlers()

	original := models.Book{
		Title:  "Test Driven Development",
		Author: "Kent Beck",
		Read:   false,
	}
	originalPayload, _ := json.Marshal(original)

	req := httptest.NewRequest(http.MethodPost, "/books", bytes.NewReader(originalPayload))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
	defer resp.Body.Close()

	var created models.Book
	err = json.NewDecoder(resp.Body).Decode(&created)
	assert.NoError(t, err)

	updated := models.Book{
		Title:  created.Title,
		Author: created.Author,
		Read:   true,
	}
	updatedPayload, _ := json.Marshal(updated)

	updateReq := httptest.NewRequest(http.MethodPut, "/books/"+strconv.Itoa(int(created.ID)), bytes.NewReader(updatedPayload))
	updateReq.Header.Set("Content-Type", "application/json")

	updateResp, err := app.Test(updateReq)
	assert.NoError(t, err)
	defer updateResp.Body.Close()
	assert.Equal(t, http.StatusOK, updateResp.StatusCode)

	var updatedBook models.Book
	err = json.NewDecoder(updateResp.Body).Decode(&updatedBook)
	assert.NoError(t, err)

	assert.Equal(t, updated.Read, updatedBook.Read)
	assert.Equal(t, created.ID, updatedBook.ID)
}

func TestDeleteBook(t *testing.T) {
	app := setupTestHandlers()

	book := models.Book{
		Title:  "Test Driven Development",
		Author: "Kent Beck",
		Read:   true,
	}
	payload, _ := json.Marshal(book)

	req := httptest.NewRequest(http.MethodPost, "/books", bytes.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
	defer resp.Body.Close()

	var created models.Book
	err = json.NewDecoder(resp.Body).Decode(&created)
	assert.NoError(t, err)

	assert.Equal(t, book.Title, created.Title)
	assert.Equal(t, book.Author, created.Author)
	assert.Equal(t, book.Read, created.Read)
	assert.NotZero(t, created.ID)

	delete_req := httptest.NewRequest(http.MethodDelete, "/books/"+strconv.Itoa(int(created.ID)), nil)
	delete_req.Header.Set("Content-Type", "application/json")

	resp, err = app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
	defer resp.Body.Close()
}
