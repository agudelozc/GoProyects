// cmd/main.go
package main

import (
	"log"
	"net/http"

	"github.com/agudelozca/go-todo-api/internal/handler"
	"github.com/agudelozca/go-todo-api/internal/middleware"
	"github.com/agudelozca/go-todo-api/internal/repository"
	"github.com/agudelozca/go-todo-api/internal/service"
	"github.com/go-chi/cors"
	"github.com/gorilla/mux"
)

func main() {
	// Initialize repository
	repo := repository.NewInMemoryTodoRepository()
	userRepo := repository.NewInMemoryUserRepository()

	// Clave para firmar los tokens
	jwtKey := []byte("my_secret")

	// Initialize service
	todoService := service.NewTodoService(repo)
	userService := service.NewUserService(userRepo, jwtKey)

	// Initialize handlers
	todoHandler := handler.NewTodoHandler(todoService)
	userHandler := handler.NewUserHandler(userService)

	// Setup router
	r := mux.NewRouter()
	r.Use(middleware.LoggingMiddleware)

	r.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           300,
	}).Handler)

	// authentication routes
	r.HandleFunc("/register", userHandler.Register).Methods("POST")
	r.HandleFunc("/login", userHandler.Login).Methods("POST")

	// api subroutes
	apiRouter := r.PathPrefix("/api").Subrouter()
	apiRouter.Use(middleware.AuthMiddleware(jwtKey))
	apiRouter.HandleFunc("/todos", todoHandler.CreateTodo).Methods("POST")
	apiRouter.HandleFunc("/todos/{id}", todoHandler.GetTodoByID).Methods("GET")
	apiRouter.HandleFunc("/todos", todoHandler.GetAllTodos).Methods("GET")
	apiRouter.HandleFunc("/todos/{id}", todoHandler.UpdateTodo).Methods("PUT")
	apiRouter.HandleFunc("/todos/{id}", todoHandler.DeleteTodoByID).Methods("DELETE")
	// admin subroutes
	adminRouter := r.PathPrefix("/admin").Subrouter()
	adminRouter.Use(middleware.AuthMiddleware(jwtKey))
	adminRouter.Use(middleware.RoleMiddleware("admin"))
	adminRouter.HandleFunc("/", handler.AdminHandler).Methods("GET")

	// Start server
	log.Println("Server started at :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
