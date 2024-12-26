package rating

import (
	"context"
	"errors"

	model "github.com/ghostriderdev/movies/rating/pkg"
)

// ErrNotFound is returned message error if record is not found.
var ErrNotFound = errors.New("not found")

type repository interface {
	Get(ctx context.Context, id model.RecordID, recordType model.RecordType) (*[]model.Rating, error)
	Put(ctx context.Context, id model.RecordID, recordType model.RecordType, rating *model.Rating) error
}

// Service defines rating service.
type Service struct {
	repo repository
}

// New creates a rating service instance.
func New(repo repository) *Service {
	return &Service{repo}
}

// GetAggregatedRating returns the aggregated rating for a
// record or ErrorNotFound if there are no ratings for it
func (s *Service) GetAggregatedRating(ctx context.Context, id model.RecordID, recordType model.RecordType) (float64, error) {
	ratings, err := s.repo.Get(ctx, id, recordType)

	if err != nil && errors.Is(err, ErrNotFound) {
		return 0, err
	} else if err != nil {
		return 0, err
	}

	sum := float64(0)

	for _, r := range *ratings {
		sum += float64(r.Value)
	}

	return sum / float64(len(*ratings)), nil
}


// PutRating writes a rating for a given record.
func (s *Service) PutRating(ctx context.Context, id model.RecordID, recordType model.RecordType, rating *model.Rating) error {
	return s.repo.Put(ctx, id, recordType, rating)
}
