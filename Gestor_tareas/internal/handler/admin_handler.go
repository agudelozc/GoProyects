// internal/handler/admin_handler.go
package handler

import (
	"net/http"

	"github.com/agudelozca/go-todo-api/internal/middleware"
	"github.com/golang-jwt/jwt"
)

func AdminHandler(w http.ResponseWriter, r *http.Request) {
	userClaims := r.Context().Value(middleware.UserContextKey).(jwt.MapClaims)
	username := userClaims["username"].(string)

	message := "Bienvenido, " + username + "! Este es un endpoint restringido para administradores."
	if _, err := w.Write([]byte(message)); err != nil {
		http.Error(w, "Error al escribir la respuesta", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
