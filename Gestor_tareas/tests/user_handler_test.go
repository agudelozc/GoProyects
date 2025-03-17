package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/agudelozca/go-todo-api/internal/handler"
	"github.com/agudelozca/go-todo-api/models"
	"github.com/agudelozca/go-todo-api/tests/mocks"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_RegisterAuthHandler(t *testing.T) {
	// arrange
	mockService := new(mocks.MockService)
	handler := handler.NewUserHandler(mockService)
	input := map[string]string{
		"username": "agudelo",
		"password": "password",
		"role":     "",
	}

	inputJSON, _ := json.Marshal(input)

	expectedUser := &models.User{
		ID:       uuid.New(),
		Username: "agudelo",
		Password: "password",
		Role:     "",
	}

	mockService.On("Register", "agudelo", "password", "").Return(expectedUser, nil)

	req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(inputJSON))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	// act
	router := chi.NewRouter()
	router.Post("/register", handler.Register)
	router.ServeHTTP(rr, req)

	// assert

	assert.Equal(t, http.StatusCreated, rr.Code)
	var resp models.User
	err = json.Unmarshal(rr.Body.Bytes(), &resp)
	assert.NoError(t, err)
	mockService.AssertExpectations(t)
}

func Test_LoginAuthHandler(t *testing.T) {
	// arrange
	mockService := new(mocks.MockService)
	handler := handler.NewUserHandler(mockService)
	input := map[string]string{
		"username": "agudelo",
		"password": "password",
	}

	inputJSON, _ := json.Marshal(input)

	expectedToken := "token"

	mockService.On("Authenticate", "agudelo", "password").Return(expectedToken, nil)

	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(inputJSON))
	assert.NoError(t, err)

	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	// act

	router := chi.NewRouter()
	router.Post("/login", handler.Login)
	router.ServeHTTP(rr, req)

	// assert

	assert.Equal(t, http.StatusOK, rr.Code)
	var resp map[string]string
	err = json.Unmarshal(rr.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.Equal(t, expectedToken, resp["token"])
	mockService.AssertExpectations(t)
}
