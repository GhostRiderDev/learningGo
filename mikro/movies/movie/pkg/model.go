package model

import model "github.com/ghostriderdev/movies/metadata/pkg/model"

// MovieDetails movie includes metadata and its aggregated rating
type MovieDetails struct {
	Rating   *float64       `json:"rating,omitempty"`
	Metadata model.Metadata `json:"metadata"`
}
