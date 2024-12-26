package rest

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	repository "github.com/ghostriderdev/movies/rating/internal/repository"
	rating "github.com/ghostriderdev/movies/rating/internal/service"
	model "github.com/ghostriderdev/movies/rating/pkg"
)

// Handler defines a rating REST handler
type Handler struct {
	service *rating.Service
}

// New crate a new rating REST handler
func New(s *rating.Service) *Handler {
	return &Handler{s}
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	recordID := model.RecordID(r.FormValue("id"))

	if recordID == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	recordType := model.RecordType(r.FormValue("type"))

	if recordType == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		v, err := h.service.GetAggregatedRating(r.Context(), recordID, recordType)

		if err != nil && errors.Is(err, repository.ErrNotFound) {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if err := json.NewEncoder(w).Encode(v); err != nil {
			log.Printf("Response encode error: %v\n", err)
		}
	case http.MethodPut:
		userId := model.UserID(r.FormValue("userId"))
		v, err := strconv.ParseFloat(r.FormValue("value"), 64)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := h.service.PutRating(r.Context(), recordID, recordType, &model.Rating{
			UserID: userId,
			Value:  model.RatingValue(v),
		}); err != nil {
			log.Printf("Repository put error: %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}
