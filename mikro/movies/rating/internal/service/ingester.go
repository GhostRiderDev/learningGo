package rating

import (
	"context"

	model "github.com/ghostriderdev/movies/rating/pkg"
)

type ingesteService interface {
	Ingest(ctx context.Context) (chan model.RatingEvent, error)
}

// StartIngestion starts the ingestion of rating events.
func (s *RatingService) StartIngestion(ctx context.Context) error {
	ch, err := s.ingester.Ingest(ctx)

	if err != nil {
		return err
	}

	for e := range ch {
		if err := s.PutRating(ctx, model.RecordID(e.RecordID), model.RecordType(e.RecordType), &model.Rating{
			Value:  e.Value,
			UserID: e.UserID,
		}); err != nil {
			return err
		}
	}
	return nil
}
