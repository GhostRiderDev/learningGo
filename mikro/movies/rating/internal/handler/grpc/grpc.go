package grpc

import (
	"context"
	"errors"

	rating "github.com/ghostriderdev/movies/rating/internal/service"
	model "github.com/ghostriderdev/movies/rating/pkg"
	"github.com/ghostriderdev/movies/src/gen"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Handler defines a grpc rating handler
type Handler struct {
	gen.UnimplementedRatingServiceServer
	service *rating.Service
}

// New creates a new instance of grpc handler rating.
func New(service *rating.Service) *Handler {
	return &Handler{service: service}
}

// GetAggregatedRating return aggregated rating for a record
func (h *Handler) GetAggregatedRating(ctx context.Context, req *gen.GetAggregatedRatingRequest) (*gen.GetAggregatedRatingResponse, error) {
	if req == nil || req.RecordId == "" || req.RecordType == "" {
		return nil, status.Errorf(codes.InvalidArgument, "req nil or recordId, recordType null")
	}

	v, err := h.service.GetAggregatedRating(ctx, model.RecordID(req.RecordId), model.RecordType(req.RecordType))

	if err != nil && errors.Is(err, rating.ErrNotFound) {
		return nil, status.Errorf(codes.NotFound, "%s", err.Error())
	} else if err != nil {
		return nil, status.Errorf(codes.Internal, "%s", err.Error())
	}
	return &gen.GetAggregatedRatingResponse{RatingValue: v}, nil
}

// PutRating writes a rating for a given record.
func (h *Handler) PutRating(ctx context.Context, req *gen.PutRatingRequest) (*gen.PutRatingResponse, error) {
	if req == nil || req.UserId == "" || req.RecordId == "" || req.RecordType == "" || req.RatingValue <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "req nil or arguments no valid")
	}

	if err := h.service.PutRating(ctx, model.RecordID(req.RecordId), model.RecordType(req.RecordType), &model.Rating{
		Value:  model.RatingValue(req.RatingValue),
		UserID: model.UserID(req.UserId),
	}); err != nil {
		return nil, err
	}
	return &gen.PutRatingResponse{}, nil
}
