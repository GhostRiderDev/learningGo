package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	metadataGateway "github.com/ghostriderdev/movies/movie/internal/gateway/metadata/rest"
	ratingGateway "github.com/ghostriderdev/movies/movie/internal/gateway/rating/rest"
	rest "github.com/ghostriderdev/movies/movie/internal/handler"
	movie "github.com/ghostriderdev/movies/movie/internal/service"
	"github.com/ghostriderdev/movies/pkg/discovery"
	"github.com/ghostriderdev/movies/pkg/discovery/consul"
)

const serviceName = "movie"

func main() {
	var port int

	flag.IntVar(&port, "port", 8080, "Api handler port")
	flag.Parse()

	log.Printf("Starting the movie service on port %d", port)

	registry, err := consul.NewRegistry("localhost:8500")

	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	instanceID := discovery.GenerateInstanceID(serviceName)

	if err := registry.Register(ctx, instanceID, serviceName, fmt.Sprintf("localhost:%d", port)); err != nil {
		panic(err)
	}

	go func() {
		for {
			if err := registry.ReportHealthyState(instanceID, serviceName); err != nil {
				log.Printf("%s", "Failed to report healthy state: "+err.Error())
			}

			time.Sleep(1 * time.Second)
		}
	}()

	defer registry.Deregister(ctx, instanceID, serviceName)



	metadataGateway := metadataGateway.New(registry)
	ratingGateway := ratingGateway.New(registry)
	ctrl := movie.New(ratingGateway, metadataGateway)
	h := rest.New(ctrl)
	http.Handle("/movie", http.HandlerFunc(h.GetMovieDetails))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
