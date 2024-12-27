package main

import (
	"context"
	"flag"
	"log"
	"net/http"

	metadataGateway "github.com/ghostriderdev/movies/movie/internal/gateway/metadata/grpc"
	ratingGateway "github.com/ghostriderdev/movies/movie/internal/gateway/rating/grpc"
	rest "github.com/ghostriderdev/movies/movie/internal/handler"
	movie "github.com/ghostriderdev/movies/movie/internal/service"
	"github.com/ghostriderdev/movies/pkg/discovery"
	"github.com/ghostriderdev/movies/pkg/discovery/consul"
)

const serviceName = "movie"

func main() {
	var port int

	flag.IntVar(&port, "port", 7070, "Api handler port")
	flag.Parse()

	log.Printf("Starting the movie service on port %d", port)

	registry, cancel, err := consul.RegisterService(port, "localhost", serviceName, "localhost:8500")
	if err != nil {
		log.Fatalf("failed to register service: %v", err)
	}
	defer cancel()
	defer registry.Deregister(context.Background(), discovery.GenerateInstanceID(serviceName), serviceName)

	metadataGateway := metadataGateway.New(registry)
	ratingGateway := ratingGateway.New(registry)
	ctrl := movie.New(ratingGateway, metadataGateway)
	h := rest.New(ctrl)
	http.Handle("/movie", http.HandlerFunc(h.GetMovieDetails))
	if err := http.ListenAndServe(":7070", nil); err != nil {
		panic(err)
	}
}
