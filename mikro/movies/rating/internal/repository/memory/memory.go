package memory

import (
	"context"

	"github.com/ghostriderdev/movies/rating/internal/repository"
	model "github.com/ghostriderdev/movies/rating/pkg"
)

// DataType define structure of data into rating repository.
type DataType map[model.RecordType]map[model.RecordID][]model.Rating

// Repository defines a rating repository.
type Repository struct {
	data DataType
}

// New create a new rating repository instance.
func New() *Repository {
	return &Repository{DataType{}}
}

// Get returns a rating by recordType and id from repository
func (r *Repository) Get(ctx context.Context, id model.RecordID, recordType model.RecordType) ([]model.Rating, error) {
	if _, ok := r.data[recordType]; !ok {
		return nil, repository.ErrNotFound
	}

	if ratings, ok := r.data[recordType][id]; !ok || len(ratings) == 0 {
		return nil, repository.ErrNotFound
	}

	return r.data[recordType][id], nil
}

// Put adds a rating for given a record
func (r *Repository) Put(ctx context.Context, id model.RecordID, recordType model.RecordType, rating *model.Rating) error {
	if _, ok := r.data[recordType]; !ok {
		r.data[recordType] = map[model.RecordID][]model.Rating{}
	}

	r.data[recordType][id] = append(r.data[recordType][id], *rating)
	return nil
}
