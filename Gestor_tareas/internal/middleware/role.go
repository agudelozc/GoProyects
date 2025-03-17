package middleware

import (
	"net/http"

	"github.com/golang-jwt/jwt"
)

// RoleMiddleware verifica que el usuario tenga el rol requerido.
func RoleMiddleware(requiredRole string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userClaims, ok := r.Context().Value(UserContextKey).(jwt.MapClaims)
			if !ok {
				http.Error(w, "No autorizado", http.StatusUnauthorized)
				return
			}

			role, ok := userClaims["role"].(string)
			if !ok || role != requiredRole {
				http.Error(w, "Permiso denegado", http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
