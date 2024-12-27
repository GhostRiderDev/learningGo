package movie

import (
	"context"
	"errors"

	model "github.com/ghostriderdev/movies/movie/pkg"
	ratingModel "github.com/ghostriderdev/movies/rating/pkg"

	metadataModel "github.com/ghostriderdev/movies/metadata/pkg/model"
)

// ErrNotFound is returned when the movie metadata is not
// found.
var ErrNotFound = errors.New("not found")

type ratingGateway interface {
	PutRating(ctx context.Context, id ratingModel.RecordID, recordType ratingModel.RecordType, rating *ratingModel.Rating) error

	GetAggregatedRating(ctx context.Context, id ratingModel.RecordID, recordType ratingModel.RecordType) (float64, error)
}

type metadataGateway interface {
	Get(ctx context.Context, id string) (*metadataModel.Metadata, error)
}

// Service defines a movie service.
type Service struct {
	ratingGateway   ratingGateway
	metadataGateway metadataGateway
}

// New crates a new instance movie service
func New(ratingGateway ratingGateway, metadataGateway metadataGateway) *Service {
	return &Service{ratingGateway, metadataGateway}
}

// Get returns the movie details including the aggregated
// rating and movie metadata.
func (s *Service) Get(ctx context.Context, id string) (*model.MovieDetails, error) {
	metadata, err := s.metadataGateway.Get(ctx, id)

	if err != nil && errors.Is(err, ErrNotFound) {
		return nil, ErrNotFound
	} else if err != nil {
		return nil, err
	}

	details := &model.MovieDetails{Metadata: *metadata}

	rating, err := s.ratingGateway.GetAggregatedRating(ctx, ratingModel.RecordID(id), ratingModel.RecordTypeMovie)

	if err != nil && errors.Is(err, ErrNotFound) {
		rating = float64(0.0)
		details.Rating = &rating
	} else if err != nil {
		return nil, err
	} else {
		details.Rating = &rating
	}
	return details, nil
}
