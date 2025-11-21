package main

import (
	"fmt"
	"net/http"
	"todoist/internal/handlers"
	"todoist/internal/repositories"
	"todoist/internal/services"
)

func main() {
	repo := repositories.NewInMemoryTodoRepo()
	service := services.NewTodoService(repo)
	handler := handlers.NewTodoHandler(service)

	http.HandleFunc("/health", handlers.HealthHandler)
	http.HandleFunc("/todos", handler.CreateTodoHandler) // POST only
	http.HandleFunc("/todos/", handler.TodoByIDHandler)  // GET, PUT, DELETE
	http.HandleFunc("/users/", handler.UsersHandler)     // GET /users/{id}/todos

	fmt.Println("server is listening on port:8080")
	http.ListenAndServe(":8080", nil)
}
