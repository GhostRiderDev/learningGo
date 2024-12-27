package metadata

import (
	"context"
	"errors"

	model "github.com/ghostriderdev/movies/metadata/pkg/model"
)

// ErrNotFound is returned when a requested record is not found
var ErrNotFound = errors.New("not found")

type metadataRepository interface {
	Get(ctx context.Context, id string) (*model.Metadata, error)
}

// Service defines metada service
type Service struct {
	repo metadataRepository
}


// New create a metadata service
func New(repo metadataRepository) *Service {
	return &Service{repo}
}

// Get returns movie metadata by id
func (c *Service) Get(ctx context.Context, id string) (*model.Metadata, error) {
	res, err := c.repo.Get(ctx, id)

	if err != nil && errors.Is(err, ErrNotFound) {
		return nil, ErrNotFound
	}

	return res, err
}
