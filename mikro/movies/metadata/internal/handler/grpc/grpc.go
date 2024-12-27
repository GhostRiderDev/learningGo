package grpc

import (
	"context"
	"errors"

	metadata "github.com/ghostriderdev/movies/metadata/internal/service"
	"github.com/ghostriderdev/movies/metadata/pkg/model"
	"github.com/ghostriderdev/movies/src/gen"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// handler defines a movie metadata grpc handler.
type Handler struct {
	gen.UnimplementedMetadataServiceServer
	service *metadata.Service
}

// New creates a new movie metadata grpc handler.
func New(service *metadata.Service) *Handler {
	return &Handler{service: service}
}

// GetMetadata returns movie metadata by id.
func (h *Handler) GetMetadata(ctx context.Context, req *gen.GetMetadataRequest) (*gen.GetMetadataResponse, error) {
	if req == nil || req.MovieId == "" {
		return nil, status.Error(codes.InvalidArgument, "nil req or empty id")
	}

	m, err := h.service.Get(ctx, req.MovieId)

	if err != nil && errors.Is(err, metadata.ErrNotFound) {
		return nil, status.Errorf(codes.NotFound, err.Error())
	} else if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &gen.GetMetadataResponse{Metadata: model.MetadataToProto(m)}, nil

}
