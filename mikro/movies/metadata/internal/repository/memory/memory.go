package memory

import (
	"context"
	"sync"

	"github.com/ghostriderdev/movies/metadata/internal/repository"
	model "github.com/ghostriderdev/movies/metadata/pkg/model"
)

// Repository defines a memory movie metadata repository
type Repository struct {
	sync.RWMutex
	data map[string]*model.Metadata
}

// New creates a new repository
func New() *Repository {
	return &Repository{data: map[string]*model.Metadata{
		"1": {
			ID:          "1",
			Title:       "Inception",
			Description: "A thief who steals corporate secrets through the use of dream-sharing technology.",
			Director:    "Christopher Nolan",
		},
		"2": {
			ID:          "2",
			Title:       "The Matrix",
			Description: "A computer hacker learns about the true nature of his reality and his role in the war against its controllers.",
			Director:    "Lana Wachowski, Lilly Wachowski",
		},
	}}
}

// Get retrieve a movie metada record by id from repository
func (r *Repository) Get(_ context.Context, id string) (*model.Metadata, error) {
	r.RLock()
	defer r.RUnlock()

	m, ok := r.data[id]

	if !ok {
		return nil, repository.ErrNotFound
	}

	return m, nil
}

// Put add or update a movie metedata record by id from the repository
func (r *Repository) Put(_ context.Context, id string, metadata *model.Metadata) error {
	r.Lock()
	defer r.Unlock()

	r.data[id] = metadata

	return nil
}
