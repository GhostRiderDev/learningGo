package rating

import kafkaingester "github.com/ghostriderdev/movies/rating/internal/ingester/kafka"

// RatingService encapsulates the rating service business
// logic
type RatingService struct {
	Service
	repo     repository
	ingester kafkaingester.Ingester
}

// New creates a rating service.
func NewRatingService(repo repository, ingester kafkaingester.Ingester) *RatingService {
	return &RatingService{repo: repo, ingester: ingester}
}
