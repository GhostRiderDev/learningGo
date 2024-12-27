package grpc

import (
	"context"

	model "github.com/ghostriderdev/movies/metadata/pkg/model"
	"github.com/ghostriderdev/movies/pkg/discovery"
	grpcutil "github.com/ghostriderdev/movies/pkg/grpc"
	"github.com/ghostriderdev/movies/src/gen"
)

// Gateway defines a movie metadata REST gateway.º
type Gateway struct {
	registry discovery.Registry
}

// New creates a new instance of metadata gateway.
func New(registry discovery.Registry) *Gateway {
	return &Gateway{registry}
}

// Get gets movie metadata by a movie id.
func (g *Gateway) Get(ctx context.Context, id string) (*model.Metadata, error) {
	conn, err := grpcutil.ServiceConnection(ctx, "metatada", g.registry)

	if err != nil {
		return nil, err
	}

	defer conn.Close()

	client := gen.NewMetadataServiceClient(conn)

	resp, err := client.GetMetadata(ctx, &gen.GetMetadataRequest{MovieId: id})

	if err != nil {
		return nil, err
	}

	return model.MetadataFromProto(resp.Metadata), nil

}
