package discovery

import (
	"context"
	"errors"
	"fmt"

	"github.com/ghostriderdev/movies/pkg/uuid"
)

// Registry defines a service registry
type Registry interface {
	Register(ctx context.Context, instanceID, serviceName, hostPort string) error
	Deregister(ctx context.Context, instanceID string, serviceName string) error

	GetInstances(ctx context.Context, serviceName string) ([]string, error)

	ReportHealthyState(instanceID, serviceName string) error
}

// ErrNotFound is returned when no instances are found for any service
var ErrNotFound = errors.New("not found")

func GenerateInstanceID(serviceName string) string {
	return fmt.Sprintf("%s-%s", serviceName, uuid.UUID())
}
