package handlers

import (
	"log"
	"net/http"

	"github.com/ghostriderdev/goTodoApi/helper"
)

func RespondWithError(w http.ResponseWriter, r *http.Request, code int, msg *string) {
	log.Println("error: ", *msg)
	ResponseWithJson(w, r, code, helper.FT{"message": *msg})
}
