package models

import "time"

type TodoStatus string

const (
	StatusPending   TodoStatus = "PENDING"
	StatusCompleted TodoStatus = "COMPLETED"
	StatusTrashed   TodoStatus = "TRASHED"
)

type Todo struct {
	ID          int
	UserID      string
	Title       string
	Description string
	Status      TodoStatus
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type CreateTodo struct {
	UserID      string
	Title       string
	Description string
}

//Update should allow partial updates, usually via pointers
type UpdateTodo struct {
	ID          int
	Title       *string
	Description *string
	Status      *TodoStatus
}
