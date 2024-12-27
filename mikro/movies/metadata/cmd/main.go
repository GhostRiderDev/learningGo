package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	grpcMetadata "github.com/ghostriderdev/movies/metadata/internal/handler/grpc"
	"github.com/ghostriderdev/movies/metadata/internal/repository/memory"
	metadata "github.com/ghostriderdev/movies/metadata/internal/service"
	"github.com/ghostriderdev/movies/pkg/discovery"
	"github.com/ghostriderdev/movies/pkg/discovery/consul"
	"github.com/ghostriderdev/movies/src/gen"
	"google.golang.org/grpc"
)

const serviceName = "metadata"

func main() {
	var port int

	flag.IntVar(&port, "port", 6060, "Api handler port")
	flag.Parse()

	log.Printf("Starting the metadata service on port %d", port)

	registry, cancel, err := consul.RegisterService(port, "localhost", serviceName, "localhost:8500")
	if err != nil {
		log.Fatalf("failed to register service: %v", err)
	}
	defer cancel()
	defer registry.Deregister(context.Background(), discovery.GenerateInstanceID(serviceName), serviceName)

	repo := memory.New()
	service := metadata.New(repo)
	h := grpcMetadata.New(service)

	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()

	gen.RegisterMetadataServiceServer(server, h)
	server.Serve(listener)
}

func implRest() {

	/*repo := memory.New()
	service := metadata.New(repo)
	h := rest.New(service)

	http.Handle("/metadata", http.HandlerFunc(h.GetMedatada))

	if err := http.ListenAndServe(fmt.Sprintf(":%d",
		port), nil); err != nil {
		panic(err)
	}
	*/
}
