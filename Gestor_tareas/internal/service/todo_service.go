package service

import (
	"errors"
	"fmt"

	"github.com/agudelozca/go-todo-api/internal/repository"
	"github.com/agudelozca/go-todo-api/models"
	"github.com/google/uuid"
)

type TodoService interface {
	CreateTodo(title string) (*models.Todo, error)
	GetTodo(id string) (*models.Todo, error)
	UpdateTodo(id string, title *string, completed *bool) (*models.Todo, error)
	DeleteTodo(id string) error
	ListTodos() ([]*models.Todo, error)
}

type todoService struct {
	repo repository.TodoRepository
}

func NewTodoService(repo repository.TodoRepository) TodoService {
	return &todoService{repo: repo}
}

func (s *todoService) CreateTodo(title string) (*models.Todo, error) {
	if title == "" {
		return nil, errors.New("el título no puede estar vacío")
	}

	todo := &models.Todo{
		ID:        uuid.New(),
		Title:     title,
		Completed: false,
	}

	if err := s.repo.SaveTodo(todo); err != nil {
		return nil, fmt.Errorf("no se pudo guardar la tarea: %w", err)
	}

	return todo, nil
}

func (s *todoService) GetTodo(id string) (*models.Todo, error) {
	if id == "" {
		return nil, errors.New("el ID no puede estar vacío")
	}

	todo, err := s.repo.FindTodoByID(id)
	if err != nil {
		return nil, err
	}

	if todo == nil {
		return nil, errors.New("tarea no encontrada")
	}

	return todo, nil
}

func (s *todoService) UpdateTodo(id string, title *string, completed *bool) (*models.Todo, error) {
	if id == "" {
		return nil, errors.New("el ID no puede estar vacío")
	}

	todo, err := s.repo.FindTodoByID(id)
	if err != nil {
		return nil, err
	}

	if todo == nil {
		return nil, errors.New("tarea no encontrada")
	}

	if title != nil {
		if *title == "" {
			return nil, errors.New("el título no puede estar vacío")
		}
		todo.Title = *title
	}

	if completed != nil {
		todo.Completed = *completed
	}

	if err := s.repo.UpdateTodo(todo); err != nil {
		return nil, fmt.Errorf("no se pudo actualizar la tarea: %w", err)
	}

	return todo, nil
}

func (s *todoService) DeleteTodo(id string) error {
	if id == "" {
		return errors.New("el ID no puede estar vacío")
	}

	err := s.repo.DeleteTodo(id)
	if err != nil {
		return fmt.Errorf("no se pudo eliminar la tarea: %w", err)
	}

	return nil
}

func (s *todoService) ListTodos() ([]*models.Todo, error) {
	todos, err := s.repo.ListTodos()
	if err != nil {
		return nil, fmt.Errorf("no se pudieron listar las tareas: %w", err)
	}

	return todos, nil
}
