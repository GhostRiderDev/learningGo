package rest

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/ghostriderdev/movies/metadata/internal/repository"
	metadata "github.com/ghostriderdev/movies/metadata/internal/service"
)

// Handler defines a movie metadata HTTP handler
type Handler struct {
	service *metadata.Service
}

// New creates a new movie metadata handler
func New(s *metadata.Service) *Handler {
	return &Handler{s}
}

func (h *Handler) GetMedatada(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	m, err := h.service.Get(ctx, id)

	if err != nil && errors.Is(err, repository.ErrNotFound) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		log.Printf("Repository Get error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(m); err != nil {
		log.Printf("Response encode error: %v\n", err)
	}

}
