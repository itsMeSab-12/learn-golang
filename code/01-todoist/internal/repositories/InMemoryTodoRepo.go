package repositories

import (
	"context"
	"sync"
	"todoist/internal/models"
)

type InMemoryTodoRepo struct {
	data   map[int]models.Todo
	autoID int
	mu     sync.RWMutex
}

func NewInMemoryTodoRepo() *InMemoryTodoRepo {
	return &InMemoryTodoRepo{
		data:   make(map[int]models.Todo),
		autoID: 1,
	}
}

// Create(ctx context.Context, t models.Todo) (models.Todo, error)
func (r *InMemoryTodoRepo) Create(ctx context.Context, t models.Todo) (models.Todo, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	t.ID = r.autoID
	r.autoID++
	r.data[t.ID] = t
	return t, nil
}

// GetByID(ctx context.Context, id int) (models.Todo, error)
func (r *InMemoryTodoRepo) GetByID(ctx context.Context, id int) (models.Todo, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	t, ok := r.data[id]
	if !ok {
		return models.Todo{}, ErrNotFound
	}

	return t, nil
}

// ListByUser(ctx context.Context, userID string) ([]models.Todo, error)
func (r *InMemoryTodoRepo) ListByUser(ctx context.Context, userID string) ([]models.Todo, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	todos := make([]models.Todo, 0, len(r.data))
	for _, t := range r.data {
		if t.UserID == userID {
			todos = append(todos, t)
		}
	}

	return todos, nil
}

// Update(ctx context.Context, id int, t models.Todo) (models.Todo, error)
func (r *InMemoryTodoRepo) Update(ctx context.Context, t models.Todo) (models.Todo, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	id := t.ID
	if _, ok := r.data[id]; !ok {
		return models.Todo{}, ErrNotFound
	}

	r.data[id] = t

	return t, nil
}

// Delete(ctx context.Context, id int) error
func (r *InMemoryTodoRepo) Delete(ctx context.Context, id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.data[id]; !ok {
		return ErrNotFound
	}
	delete(r.data, id)
	return nil
}
