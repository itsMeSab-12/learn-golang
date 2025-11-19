package services

import (
	"context"
	"time"
	"todoist/internal/models"
	"todoist/internal/repositories"
)

type ITodoService interface {
	CreateTodo(ctx context.Context, dto models.CreateTodo) (models.Todo, error)
	GetTodo(ctx context.Context, id int) (models.Todo, error)
	ListTodos(ctx context.Context, userID string) ([]models.Todo, error)
	UpdateTodo(ctx context.Context, id int, dto models.UpdateTodo) (models.Todo, error)
	DeleteTodo(ctx context.Context, id int) error
}

type TodoService struct {
	repo repositories.TodoRepository
}

func NewTodoService(repo repositories.TodoRepository) *TodoService {
	return &TodoService{
		repo: repo,
	}
}

// CreateTodo validates input, constructs domain model, and delegates to repository
func (s *TodoService) CreateTodo(ctx context.Context, dto models.CreateTodo) (models.Todo, error) {

	if dto.UserID == "" || dto.Title == "" {
		return models.Todo{}, ErrInvalidInput
	}

	if len(dto.Title) > 255 {
		return models.Todo{}, ErrInvalidInput
	}

	t := models.Todo{
		UserID:      dto.UserID,
		Title:       dto.Title,
		Description: dto.Description,
		Status:      models.StatusPending,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	return s.repo.Create(ctx, t)
}

// GetTodo retrieves a todo by ID with validation
func (s *TodoService) GetTodo(ctx context.Context, id int) (models.Todo, error) {
	if id <= 0 {
		return models.Todo{}, ErrInvalidInput
	}

	return s.repo.GetByID(ctx, id)
}

// ListTodos retrieves all todos for a given user
func (s *TodoService) ListTodos(ctx context.Context, userID string) ([]models.Todo, error) {
	if userID == "" {
		return nil, ErrInvalidInput
	}

	return s.repo.ListByUser(ctx, userID)
}

func (s *TodoService) UpdateTodo(ctx context.Context, dto models.UpdateTodo) (models.Todo, error) {

	if id := dto.ID; id <= 0 {
		return models.Todo{}, ErrInvalidInput
	}

	existing, err := s.repo.GetByID(ctx, dto.ID)

	if err != nil {
		// propagate ErrNotFound from repo
		return models.Todo{}, err
	}

	if dto.Title != nil {
		if *dto.Title == "" || len(*dto.Title) > 255 {
			return models.Todo{}, ErrInvalidInput
		}
		existing.Title = *dto.Title
	}
	if dto.Description != nil {
		existing.Description = *dto.Description
	}
	if dto.Status != nil {
		switch *dto.Status {
		case models.StatusPending, models.StatusCompleted, models.StatusTrashed:
			existing.Status = *dto.Status
		default:
			return models.Todo{}, ErrInvalidInput
		}
	}

	existing.UpdatedAt = time.Now()

	return s.repo.Update(ctx, existing)
}

func (s *TodoService) DeleteTodo(ctx context.Context, id int) error {
	if id <= 0 {
		return ErrInvalidInput
	}

	return s.repo.Delete(ctx, id)
}
