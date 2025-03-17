package handler

import (
	"encoding/json"
	"net/http"

	"github.com/agudelozca/go-todo-api/internal/service"
)

// AuthHandler handles HTTP requests for Users.
type AuthHandler struct {
	service service.UserService
}

// NewUserHandler creates a new UserHandler.
func NewUserHandler(service service.UserService) *AuthHandler {
	return &AuthHandler{service: service}
}

type RegisterResponse struct {
	Username string `json:"username"`
	Role     string `json:"role"`
}

// Register maneja el registro de nuevos usuarios.
func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Role     string `json:"role"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Entrada inválida", http.StatusBadRequest)
		return
	}

	user, err := h.service.Register(input.Username, input.Password, input.Role)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	RegisterResponse := RegisterResponse{
		Username: user.Username,
		Role:     user.Role,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(RegisterResponse); err != nil {
		http.Error(w, "Error al generar la respuesta", http.StatusInternalServerError)
	}
}

// Login maneja la autenticación de usuarios.
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Entrada inválida", http.StatusBadRequest)
		return
	}

	token, err := h.service.Authenticate(input.Username, input.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	response := map[string]string{
		"token": token,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error al generar la respuesta", http.StatusInternalServerError)
	}
}
