package handlers

import (
	"net/http"

	"github.com/ghostriderdev/goTodoApi/helper"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		message := "Method no allowed"
		RespondWithError(w, r, http.StatusMethodNotAllowed, &message)
		return
	}

	ResponseWithJson(w, r, http.StatusOK, helper.FT{"message": "OK"})
}
