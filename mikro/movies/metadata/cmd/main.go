package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

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

	registryService(port)

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
	service := metadata.New(repo)
	h := rest.New(service)

	http.Handle("/metadata", http.HandlerFunc(h.GetMedatada))

	if err := http.ListenAndServe(fmt.Sprintf(":%d",
		port), nil); err != nil {
		panic(err)
	}
	*/
}
