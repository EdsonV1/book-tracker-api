package handlers

import (
	"book-tracker-api/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestRegisterUser(t *testing.T) {
	app := setupTestHandlers()
	user := models.User{
		Username: "user23",
		Password: "1234",
	}

	payload, _ := json.Marshal(user)

	req := httptest.NewRequest(http.MethodPost, "/auth/register", bytes.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
	defer resp.Body.Close()
}

func TestLoginUser(t *testing.T) {
	app := setupTestHandlers()
	user := models.User{
		Username: "user",
		Password: "1234",
	}

	payload, _ := json.Marshal(user)

	req := httptest.NewRequest(http.MethodPost, "/auth/register", bytes.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusCreated, resp.StatusCode)

	defer resp.Body.Close()

	login_user := models.User{
		Username: "user",
		Password: "1234",
	}

	payload, _ = json.Marshal(login_user)

	req_login := httptest.NewRequest(http.MethodPost, "/auth/login", bytes.NewReader(payload))
	req_login.Header.Set("Content-Type", "application/json")

	resp, err = app.Test(req_login)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
}
