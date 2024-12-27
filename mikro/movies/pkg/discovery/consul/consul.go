package consul

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/ghostriderdev/movies/pkg/discovery"
	consul "github.com/hashicorp/consul/api"
)

// Registry defines Consul-based registry service.
type Registry struct {
	client *consul.Client
}

// NewRegistry create a new Consul-based registry service.
func NewRegistry(addr string) (*Registry, error) {
	config := consul.DefaultConfig()
	config.Address = addr

	client, err := consul.NewClient(config)

	if err != nil {
		return nil, err
	}

	return &Registry{client}, nil
}

// Register creates a service record in the registry.
func (r *Registry) Register(ctx context.Context, instanceID, serviceName, hostPort string) error {
	parts := strings.Split(hostPort, ":")

	if len(parts) != 2 {
		return errors.New("hostPort must be in form of <host>:<port>, example: localhost:8081")
	}

	port, err := strconv.Atoi(parts[1])

	if err != nil {
		return err
	}

	return r.client.Agent().ServiceRegister(&consul.AgentServiceRegistration{
		Address: parts[0],
		Port:    port,
		ID:      instanceID,
		Name:    serviceName,
		Check:   &consul.AgentServiceCheck{CheckID: instanceID, TTL: "5s"},
	})
}

// Deregister removes a service record from the
// registry.
func (r *Registry) Deregister(ctx context.Context, instanceID string, _ string) error {
	return r.client.Agent().ServiceDeregister(instanceID)
}

// ServiceAddresses returns the list of addresses of
// active instances of the given service
func (r *Registry) GetInstances(ctx context.Context, serviceName string) ([]string, error) {
	entries, _, err := r.client.Health().Service(serviceName, "", true, nil)

	if err != nil {
		return nil, err
	} else if len(entries) == 0 {
		return nil, discovery.ErrNotFound
	}
	var res []string

	for _, e := range entries {
		res = append(res, fmt.Sprintf("%s:%d", e.Service.Address, e.Service.Port))
	}
	return res, nil
}

func (r *Registry) ReportHealthyState(instanceID, _ string) error {
	return r.client.Agent().UpdateTTL(instanceID, "Service is healthy", "passing")
}
