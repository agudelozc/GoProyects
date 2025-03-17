package mocks

import (
	"github.com/agudelozca/go-todo-api/models"
	"github.com/stretchr/testify/mock"
)

// MockService is a mock implementation of the TodoService interface

type MockService struct {
	mock.Mock
}

func (m *MockService) CreateTodo(title string) (*models.Todo, error) {
	args := m.Called(title)
	return args.Get(0).(*models.Todo), args.Error(1)
}

func (m *MockService) GetTodo(id string) (*models.Todo, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Todo), args.Error(1)
}

func (m *MockService) ListTodos() ([]*models.Todo, error) {
	args := m.Called()
	return args.Get(0).([]*models.Todo), args.Error(1)
}

func (m *MockService) UpdateTodo(id string, title *string, completed *bool) (*models.Todo, error) {
	args := m.Called(id, title, completed)
	if args.Get(0) != nil {
		return args.Get(0).(*models.Todo), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockService) DeleteTodo(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

// MockRepository is a mock implementation of the TodoRepository interface

type MockRepository struct {
	mock.Mock
}

// FindTodoByID implements repository.TodoRepository.
func (m *MockRepository) FindTodoByID(id string) (*models.Todo, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Todo), args.Error(1)
}

// ListTodos implements repository.TodoRepository.
func (m *MockRepository) ListTodos() ([]*models.Todo, error) {
	args := m.Called()
	return args.Get(0).([]*models.Todo), args.Error(1)
}

// SaveTodo implements repository.TodoRepository.
func (m *MockRepository) SaveTodo(todo *models.Todo) error {
	args := m.Called(todo)
	return args.Error(0)
}

// UpdateTodo implements repository.TodoRepository.
func (m *MockRepository) UpdateTodo(todo *models.Todo) error {
	args := m.Called(todo)
	return args.Error(0)
}

// DeleteTodo implements repository.TodoRepository.
func (m *MockRepository) DeleteTodo(id string) error {
	args := m.Called(id)
	return args.Error(0)
}
