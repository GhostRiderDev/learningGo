package main

import (
	"log"
	"net/http"

	rest "github.com/ghostriderdev/movies/metadata/internal/handler"
	"github.com/ghostriderdev/movies/metadata/internal/repository/memory"
	metadata "github.com/ghostriderdev/movies/metadata/internal/service"
)

func main() {
	log.Println("Starting movie metadata service")
	repo := memory.New()
	service := metadata.New(repo)
	h := rest.New(service)

	http.Handle("/metadata", http.HandlerFunc(h.GetMedatada))

	if err := http.ListenAndServe(":6060", nil); err != nil {
		panic(err)
	}
}
