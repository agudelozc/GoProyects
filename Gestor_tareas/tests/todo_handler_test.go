package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/agudelozca/go-todo-api/internal/handler"
	"github.com/agudelozca/go-todo-api/models"
	"github.com/agudelozca/go-todo-api/tests/mocks"
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/mux"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateTodoHandler(t *testing.T) {
	// arrange
	mockService := new(mocks.MockService)
	handler := handler.NewTodoHandler(mockService)

	input := map[string]string{
		"title": "Learn TDD with Go",
	}
	inputJSON, _ := json.Marshal(input)

	expectedTodo := &models.Todo{
		ID:        uuid.New(),
		Title:     "Learn TDD with Go",
		Completed: false,
	}

	mockService.On("CreateTodo", "Learn TDD with Go").Return(expectedTodo, nil)

	req, err := http.NewRequest("POST", "/todos", bytes.NewBuffer(inputJSON))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	// act
	router := chi.NewRouter()
	router.Post("/todos", handler.CreateTodo)
	router.ServeHTTP(rr, req)

	// assert
	assert.Equal(t, http.StatusCreated, rr.Code)

	var resp models.Todo
	err = json.Unmarshal(rr.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.Equal(t, expectedTodo.Title, resp.Title)
	assert.Equal(t, expectedTodo.Completed, resp.Completed)
	assert.Equal(t, expectedTodo.ID, resp.ID)

	mockService.AssertExpectations(t)
}

func TestGetTodoByID(t *testing.T) {
	// arrange
	mockService := new(mocks.MockService)
	handler := handler.NewTodoHandler(mockService)

	testUUID := uuid.MustParse("00000000-0000-0000-0000-000000000001")
	expectedTodo := &models.Todo{
		ID:        testUUID,
		Title:     "Learn TDD with Go",
		Completed: false,
	}

	mockService.On("GetTodo", testUUID.String()).Return(expectedTodo, nil)

	// Create a request with the test UUID
	req, err := http.NewRequest("GET", "/todos/"+testUUID.String(), nil)
	assert.NoError(t, err)

	// Use gomux to set the route variables
	req = mux.SetURLVars(req, map[string]string{
		"id": testUUID.String(),
	})

	// Create a ResponseRecorder to capture the response
	rr := httptest.NewRecorder()

	// Call the handler
	handler.GetTodoByID(rr, req)

	// Assertions
	assert.Equal(t, http.StatusOK, rr.Code)
	var todoResp models.Todo
	err = json.Unmarshal(rr.Body.Bytes(), &todoResp)
	assert.NoError(t, err)
	assert.Equal(t, expectedTodo, &todoResp)

	// Assert that the expectations were met
	mockService.AssertExpectations(t)
}

func TestGetAllTodos(t *testing.T) {
	// arrange
	mockService := new(mocks.MockService)
	handler := handler.NewTodoHandler(mockService)

	expectedTodos := []*models.Todo{
		{
			ID:        uuid.MustParse("00000000-0000-0000-0000-000000000001"),
			Title:     "Learn TDD with Go",
			Completed: false,
		},
		{
			ID:        uuid.MustParse("00000000-0000-0000-0000-000000000002"),
			Title:     "Learn TDD with Python",
			Completed: false,
		},
	}

	mockService.On("ListTodos").Return(expectedTodos, nil)

	req := httptest.NewRequest(http.MethodGet, "/todos", nil)
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	// act
	router := chi.NewRouter()
	router.Get("/todos", handler.GetAllTodos)
	router.ServeHTTP(rr, req)

	// assert
	expectedBody := `[{"id":"00000000-0000-0000-0000-000000000001","title":"Learn TDD with Go","completed":false},{"id":"00000000-0000-0000-0000-000000000002","title":"Learn TDD with Python","completed":false}]`
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, expectedBody, strings.TrimSpace(rr.Body.String()))
	mockService.AssertExpectations(t)
}

func TestUpdateTodo(t *testing.T) {
	// arrange
	mockService := new(mocks.MockService)
	handler := handler.NewTodoHandler(mockService)

	testUUID := uuid.MustParse("00000000-0000-0000-0000-000000000001")
	title := "Learn TDD with Go"
	completed := true
	expectedTodo := &models.Todo{
		ID:        testUUID,
		Title:     title,
		Completed: completed,
	}

	mockService.On("UpdateTodo", testUUID.String(), &title, &completed).Return(expectedTodo, nil)

	reqBody, err := json.Marshal(map[string]interface{}{
		"title":     title,
		"completed": completed,
	})
	assert.NoError(t, err)

	// Create an HTTP PUT request with the ID and payload
	req, err := http.NewRequest("PUT", "/todos/"+testUUID.String(), bytes.NewBuffer(reqBody))
	assert.NoError(t, err)

	// Set the route variables using mux
	req = mux.SetURLVars(req, map[string]string{
		"id": testUUID.String(),
	})

	// Set the content type
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	// act
	router := chi.NewRouter()
	router.Put("/todos/{id}", handler.UpdateTodo)
	router.ServeHTTP(rr, req)

	// assert
	assert.Equal(t, http.StatusOK, rr.Code)
	var todoResp models.Todo
	err = json.Unmarshal(rr.Body.Bytes(), &todoResp)
	assert.NoError(t, err)
	assert.Equal(t, expectedTodo, &todoResp)

	mockService.AssertExpectations(t)
}

func TestDeleteTodoByID(t *testing.T) {
	// arrange
	mockService := new(mocks.MockService)
	handler := handler.NewTodoHandler(mockService)
	testUUID := uuid.MustParse("00000000-0000-0000-0000-000000000001")
	mockService.On("DeleteTodo", testUUID.String()).Return(nil)

	// Create an HTTP PUT request with the ID and payload
	req, err := http.NewRequest("DELETE", "/todos/"+testUUID.String(), nil)
	assert.NoError(t, err)

	// Set the route variables using mux
	req = mux.SetURLVars(req, map[string]string{
		"id": testUUID.String(),
	})

	// Set the content type
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	// act
	router := chi.NewRouter()
	router.Delete("/todos/{id}", handler.DeleteTodoByID)
	router.ServeHTTP(rr, req)

	// assert
	assert.Equal(t, http.StatusNoContent, rr.Code)
	mockService.AssertExpectations(t)
}

func TestCreateTodoHandler_InvalidInput(t *testing.T) {
	// arrange
	mockService := new(mocks.MockService)
	handler := handler.NewTodoHandler(mockService)
	mockService.On("CreateTodo", "").Return(&models.Todo{}, errors.New("title cannot be empty"))

	input := map[string]string{
		"title": "",
	}
	inputJSON, _ := json.Marshal(input)

	req, err := http.NewRequest("POST", "/todos", bytes.NewBuffer(inputJSON))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	// act
	router := chi.NewRouter()
	router.Post("/todos", handler.CreateTodo)
	router.ServeHTTP(rr, req)

	// assert
	assert.Equal(t, http.StatusBadRequest, rr.Code)
	mockService.AssertExpectations(t)
}

func TestUpdateTodoByID_InvalidID(t *testing.T) {
	// arrange
	mockService := new(mocks.MockService)
	handler := handler.NewTodoHandler(mockService)

	req := httptest.NewRequest(http.MethodPut, "/todos/invalid-id", nil)
	chictx := chi.NewRouteContext()
	chictx.URLParams.Add("id", "invalid-id")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chictx))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	// act
	router := chi.NewRouter()
	router.Put("/todos/{id}", handler.UpdateTodo)
	router.ServeHTTP(rr, req)

	// assert
	assert.Equal(t, http.StatusBadRequest, rr.Code)
	mockService.AssertExpectations(t)
}
