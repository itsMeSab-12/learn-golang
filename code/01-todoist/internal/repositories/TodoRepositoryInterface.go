package repositories

import (
	"context"
	"todoist/internal/models"
)

type TodoRepository interface {
	Create(ctx context.Context, t models.Todo) (models.Todo, error)
	GetByID(ctx context.Context, id int) (models.Todo, error)
	ListByUser(ctx context.Context, userID string) ([]models.Todo, error)
	Update(ctx context.Context, t models.Todo) (models.Todo, error)
	Delete(ctx context.Context, id int) error
}
