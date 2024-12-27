package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"

	"github.com/ghostriderdev/movies/movie/internal/gateway"
	"github.com/ghostriderdev/movies/pkg/discovery"
	model "github.com/ghostriderdev/movies/rating/pkg"
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
	addrs, err := g.registry.GetInstances(ctx, "rating")

	if err != nil {
		return 0, err
	}

	url := "http://" + addrs[rand.Intn(len(addrs))] + "/rating"
	log.Printf("%s", "Calling metadata service. Request: GET "+url)

	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return 0, err
	}

	req = req.WithContext(ctx)
	values := req.URL.Query()
	values.Add("id", string(id))
	values.Add("type", string(recordType))
	req.URL.RawQuery = values.Encode()

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return 0, gateway.ErrNotFound
	} else if resp.StatusCode/100 != 2 {
		return 0, fmt.Errorf("non-2xx response: %v", resp)
	}

	var v float64

	if err := json.NewDecoder(resp.Body).Decode(&v); err != nil {
		return 0, err
	}

	return v, nil
}

// PutRating writes a rating.
func (g *Gateway) PutRating(ctx context.Context, id model.RecordID, recordType model.RecordType, rating *model.Rating) error {
	addrs, err := g.registry.GetInstances(ctx, "metadata")

	if err != nil {
		return err
	}

	url := "http://" + addrs[rand.Intn(len(addrs))] + "/metadata"
	log.Printf("%s", "Calling metadata service. Request: GET "+url)

	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return err
	}

	req = req.WithContext(ctx)

	values := req.URL.Query()
	values.Add("id", string(id))
	values.Add("type", string(recordType))
	values.Add("userId", string(rating.UserID))
	values.Add("value", string(rating.Value))

	req.URL.RawQuery = values.Encode()

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode/100 != 2 {
		return fmt.Errorf("non-2xx response: %v", resp)
	}
	return nil
}
