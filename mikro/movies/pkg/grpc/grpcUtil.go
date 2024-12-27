package grpcutil

import (
	"context"
	"math/rand"

	"github.com/ghostriderdev/movies/pkg/discovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ServiceConnection(ctx context.Context, serviceName string, registry discovery.Registry) (*grpc.ClientConn, error) {
	addrs, err := registry.GetInstances(ctx, serviceName)

	if err != nil {
		return nil, err
	}

	return grpc.NewClient(addrs[rand.Intn(len(addrs))], grpc.WithTransportCredentials(insecure.NewCredentials()))
}
