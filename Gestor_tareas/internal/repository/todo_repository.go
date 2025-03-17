// internal/repository/todo_repository.go
package repository

import (
	"errors"

	"github.com/agudelozca/go-todo-api/models"
)

// TodoRepository define los métodos para la persistencia de tareas.
type TodoRepository interface {
	SaveTodo(todo *models.Todo) error
	FindTodoByID(id string) (*models.Todo, error)
	UpdateTodo(todo *models.Todo) error
	DeleteTodo(id string) error
	ListTodos() ([]*models.Todo, error)
}

// InMemoryTodoRepository es una implementación en memoria de TodoRepository.
type InMemoryTodoRepository struct {
	todos map[string]*models.Todo
}

// NewInMemoryTodoRepository crea una nueva instancia de InMemoryTodoRepository.
func NewInMemoryTodoRepository() *InMemoryTodoRepository {
	return &InMemoryTodoRepository{
		todos: make(map[string]*models.Todo),
	}
}

// SaveTodo guarda una tarea en memoria.
func (r *InMemoryTodoRepository) SaveTodo(todo *models.Todo) error {
	if _, exists := r.todos[todo.ID.String()]; exists {
		return errors.New("la tarea ya existe")
	}
	r.todos[todo.ID.String()] = todo
	return nil
}

// FindTodoByID encuentra una tarea por su ID.
func (r *InMemoryTodoRepository) FindTodoByID(id string) (*models.Todo, error) {
	todo, exists := r.todos[id]
	if !exists {
		return nil, nil
	}
	return todo, nil
}

// UpdateTodo actualiza una tarea existente en memoria.
func (r *InMemoryTodoRepository) UpdateTodo(todo *models.Todo) error {
	if _, exists := r.todos[todo.ID.String()]; !exists {
		return errors.New("la tarea no existe")
	}
	r.todos[todo.ID.String()] = todo
	return nil
}

// DeleteTodo elimina una tarea por su ID.
func (r *InMemoryTodoRepository) DeleteTodo(id string) error {
	if _, exists := r.todos[id]; !exists {
		return errors.New("la tarea no existe")
	}
	delete(r.todos, id)
	return nil
}

// ListTodos lista todas las tareas.
func (r *InMemoryTodoRepository) ListTodos() ([]*models.Todo, error) {
	var list []*models.Todo
	for _, todo := range r.todos {
		list = append(list, todo)
	}
	return list, nil
}
