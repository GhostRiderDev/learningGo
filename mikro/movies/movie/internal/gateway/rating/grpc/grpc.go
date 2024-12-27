package grpc

import (
	"context"

	"github.com/ghostriderdev/movies/pkg/discovery"
	grpcutil "github.com/ghostriderdev/movies/pkg/grpc"
	model "github.com/ghostriderdev/movies/rating/pkg"
	"github.com/ghostriderdev/movies/src/gen"
)

// Gateway defines an REST gateway for a rating service.
type Gateway struct {
	registry discovery.Registry
}

// New creates a new instance for a rating REST gateway service.
func New(registry discovery.Registry) *Gateway {
	return &Gateway{registry}
}

// GetAggregatedRating returns the aggregated rating for a record
// or ErrNotFound if there are no ratings for it.
func (g *Gateway) GetAggregatedRating(ctx context.Context, id model.RecordID, recordType model.RecordType) (float64, error) {
	conn, err := grpcutil.ServiceConnection(ctx, "rating", g.registry)

	if err != nil {
		return 0, err
	}

	defer conn.Close()

	client := gen.NewRatingServiceClient(conn)

	resp, err := client.GetAggregatedRating(ctx, &gen.GetAggregatedRatingRequest{
		RecordId:   string(id),
		RecordType: string(recordType),
	})

	if err != nil {
		return 0, err
	}

	return resp.RatingValue, nil
}

// PutRating writes a rating.
func (g *Gateway) PutRating(ctx context.Context, id model.RecordID, recordType model.RecordType, rating *model.Rating) error {
	conn, err := grpcutil.ServiceConnection(ctx, "rating", g.registry)

	if err != nil {
		return err
	}

	defer conn.Close()

	client := gen.NewRatingServiceClient(conn)

	_, err = client.PutRating(ctx, &gen.PutRatingRequest{
		RecordId:    string(id),
		RecordType:  string(recordType),
		UserId:      string(rating.UserID),
		RatingValue: float64(rating.Value),
	})
	if err != nil {
		return err
	}

	return nil
}
