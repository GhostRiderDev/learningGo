package rest

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	movie "github.com/ghostriderdev/movies/movie/internal/service"
)

// Handler defines a movie handler.
type Handler struct {
	service *movie.Service
}

// New creates a new movie HTTP handler.
func New(service *movie.Service) *Handler {
	return &Handler{service}
}

// GetMovieDetails handles GET /movie requests.
func (h *Handler) GetMovieDetails(w http.ResponseWriter, req *http.Request) {
	id := req.FormValue("id")

	w.Header().Add("Content-Type", "application/json")

	details, err := h.service.Get(req.Context(), id)
	if err != nil && errors.Is(err, movie.ErrNotFound) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		log.Printf("Repository get error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(details); err != nil {
		log.Printf("Response encode error: %v\n", err)
	}
}
