package models

import "time"

type TodoStatus string

const (
	StatusPending   TodoStatus = "PENDING"
	StatusCompleted TodoStatus = "COMPLETED"
	StatusTrashed   TodoStatus = "TRASHED"
)

type Todo struct {
	ID          int        `json:"id"`
	UserID      string     `json:"userid"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      TodoStatus `json:"status"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
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
