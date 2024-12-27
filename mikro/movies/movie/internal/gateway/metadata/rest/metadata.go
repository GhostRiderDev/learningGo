package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"

	model "github.com/ghostriderdev/movies/metadata/pkg/model"
	"github.com/ghostriderdev/movies/movie/internal/gateway"
	"github.com/ghostriderdev/movies/pkg/discovery"
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
	addrs, err := g.registry.GetInstances(ctx, "metadata")

	if err != nil {
		return nil, err
	}

	url := "http://" + addrs[rand.Intn(len(addrs))] + "/metadata"
	log.Printf("%s", "Calling metadata service. Request: GET " + url)

	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	values := req.URL.Query()
	values.Add("id", id)
	req.URL.RawQuery = values.Encode()

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, gateway.ErrNotFound
	} else if resp.StatusCode/100 != 2 {
		return nil, fmt.Errorf("non-2xx response: %v", resp)
	}

	var v *model.Metadata

	if err := json.NewDecoder(resp.Body).Decode(&v); err != nil {
		log.Fatalf("Error marshaling body at gateway: %v", err)
		return nil, err
	}

	return v, nil
}
