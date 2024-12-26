package main

import (
	"net/http"

	rest "github.com/ghostriderdev/movies/rating/internal/handler"
	"github.com/ghostriderdev/movies/rating/internal/repository/memory"
	rating "github.com/ghostriderdev/movies/rating/internal/service"
)

func main() {
	repo := memory.New()
	service := rating.New(repo)
	h := rest.New(service)

	http.Handle("/rating", http.HandlerFunc(h.Handle))

	if err := http.ListenAndServe(":5050", nil); err != nil {
		panic(err)
	}
}
