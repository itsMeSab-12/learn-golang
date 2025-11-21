package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"todoist/internal/models"
	"todoist/internal/services"
)

type TodoHandler struct {
	Service services.ITodoService
}

func NewTodoHandler(s services.ITodoService) *TodoHandler {
	return &TodoHandler{Service: s}
}

func (h *TodoHandler) TodoByIDHandler(w http.ResponseWriter, r *http.Request) {
	if !strings.HasPrefix(r.URL.Path, "/todos/") {
		http.NotFound(w, r)
		return
	}

	idStr := strings.TrimPrefix(r.URL.Path, "/todos/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		h.GetTodo(w, r, id)
	case http.MethodPut:
		h.UpdateTodo(w, r, id)
	case http.MethodDelete:
		h.DeleteTodo(w, r, id)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}

}

func (h *TodoHandler) UsersHandler(w http.ResponseWriter, r *http.Request) {
	if !strings.HasPrefix(r.URL.Path, "/users/") {
		http.NotFound(w, r)
		return
	}

	parts := strings.Split(r.URL.Path, "/")
	// /users/{id}/todos -> ["","users","{id}", "todos"]
	if len(parts) != 4 || parts[3] != "todos" {
		http.NotFound(w, r)
		return
	}

	userID := parts[2]

	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	todos, err := h.Service.ListTodos(r.Context(), userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(todos)
}

func (h *TodoHandler) CreateTodoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		dto := models.CreateTodo{}

		if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
			http.Error(w, "invalid JSON", http.StatusBadRequest)
			return
		}

		todo, err := h.Service.CreateTodo(r.Context(), dto)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(todo)
		return
	}

	http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
}

func (h *TodoHandler) GetTodo(w http.ResponseWriter, r *http.Request, id int) {
	todo, err := h.Service.GetTodo(r.Context(), id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(todo)
}

func (h *TodoHandler) UpdateTodo(w http.ResponseWriter, r *http.Request, id int) {
	dto := models.UpdateTodo{}

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	dto.ID = id

	todo, err := h.Service.UpdateTodo(r.Context(), dto)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(todo)
}

func (h *TodoHandler) DeleteTodo(w http.ResponseWriter, r *http.Request, id int) {
	if err := h.Service.DeleteTodo(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
