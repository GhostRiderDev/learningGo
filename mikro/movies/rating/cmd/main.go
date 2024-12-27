package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/ghostriderdev/movies/pkg/discovery"
	"github.com/ghostriderdev/movies/pkg/discovery/consul"
	grpcRating "github.com/ghostriderdev/movies/rating/internal/handler/grpc"
	"github.com/ghostriderdev/movies/rating/internal/repository/memory"
	rating "github.com/ghostriderdev/movies/rating/internal/service"
	"github.com/ghostriderdev/movies/src/gen"
	"google.golang.org/grpc"
)

const serviceName = "rating"

func main() {
	var port int

	flag.IntVar(&port, "port", 5050, "Api handler port")
	flag.Parse()

	log.Printf("Starting the rating service on port %d", port)

	registry, cancel, err := consul.RegisterService(port, "localhost", serviceName, "localhost:8500")
	if err != nil {
		log.Fatalf("failed to register service: %v", err)
	}
	defer cancel()
	defer registry.Deregister(context.Background(), discovery.GenerateInstanceID(serviceName), serviceName)

	repo := memory.New()
	service := rating.New(repo)
	h := grpcRating.New(service)

	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))

	if err != nil {
		log.Fatalf("failed to listen: %v", err.Error())
	}

	server := grpc.NewServer()
	gen.RegisterRatingServiceServer(server, h)
	server.Serve(listener)

}

func implRest() {
	/*repo := memory.New()
	service := rating.New(repo)
	h := rest.New(service)

	http.Handle("/rating", http.HandlerFunc(h.Handle))

	if err := http.ListenAndServe(fmt.Sprintf(":%d",
		port), nil); err != nil {
		panic(err)
	}
	*/
}
