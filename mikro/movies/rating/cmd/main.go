package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

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

	flag.IntVar(&port, "port", 6060, "Api handler port")
	flag.Parse()

	log.Printf("Starting the rating service on port %d", port)

	registryService(port)

	repo := memory.New()
	service := rating.New(repo)
	h := grpcRating.New(service)

	listener, err := net.Listen("tcp", "localhost:6060")

	if err != nil {
		log.Fatalf("failed to listen: %v", err.Error())
	}

	server := grpc.NewServer()
	gen.RegisterRatingServiceServer(server, h)
	server.Serve(listener)

}

func registryService(port int) {
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
