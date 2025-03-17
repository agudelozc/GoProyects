package handler

import (
	"encoding/json"
	"net/http"

	"github.com/agudelozca/go-todo-api/internal/service"
	"github.com/gorilla/mux"
)

// TodoHandler handles HTTP requests for todos.
type TodoHandler struct {
	service service.TodoService
}

// NewTodoHandler creates a new TodoHandler.
func NewTodoHandler(service service.TodoService) *TodoHandler {
	return &TodoHandler{service: service}
}

// CreateTodo handles the creation of a new todo.
func (h *TodoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title string `json:"title"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	todo, err := h.service.CreateTodo(input.Title)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	respondWithJSON(w, http.StatusCreated, todo)
}

// GetTodoByID handles the retrieval of a todo by its ID.
func (h *TodoHandler) GetTodoByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	todo, err := h.service.GetTodo(id)
	if err != nil {
		status := http.StatusBadRequest
		if err.Error() == "tarea no encontrada" {
			status = http.StatusNotFound
		}
		http.Error(w, err.Error(), status)
		return
	}

	respondWithJSON(w, http.StatusOK, todo)
}

// GetAllTodos handles the retrieval of all todos.
func (h *TodoHandler) GetAllTodos(w http.ResponseWriter, _ *http.Request) {
	todos, err := h.service.ListTodos()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, todos)
}

// UpdateTodo handles the update of an existing todo.
func (h *TodoHandler) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var input struct {
		Title     *string `json:"title,omitempty"`
		Completed *bool   `json:"completed,omitempty"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Entrada inválida", http.StatusBadRequest)
		return
	}

	todo, err := h.service.UpdateTodo(id, input.Title, input.Completed)
	if err != nil {
		status := http.StatusBadRequest
		if err.Error() == "tarea no encontrada" {
			status = http.StatusNotFound
		}
		http.Error(w, err.Error(), status)
		return
	}

	respondWithJSON(w, http.StatusOK, todo)
}

// DeleteTodoByID handles the deletion of a todo by its ID.
func (h *TodoHandler) DeleteTodoByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	if err := h.service.DeleteTodo(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// respondWithJSON writes a JSON response to the ResponseWriter.
func respondWithJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
