package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	grpcMetadata "github.com/ghostriderdev/movies/metadata/internal/handler/grpc"
	dbmysql "github.com/ghostriderdev/movies/metadata/internal/repository/mysql"
	metadata "github.com/ghostriderdev/movies/metadata/internal/service"
	"github.com/ghostriderdev/movies/pkg/discovery"
	"github.com/ghostriderdev/movies/pkg/discovery/consul"
	"github.com/ghostriderdev/movies/src/gen"
	"google.golang.org/grpc"
	"gopkg.in/yaml.v3"
)

const serviceName = "metadata"

func main() {

	log.Printf("Starting the metadata service on port")

	f, err := os.Open("base.yaml")

	if err != nil {
		panic(err)
	}
	defer f.Close()

	var cfg serviceConfig

	if err := yaml.NewDecoder(f).Decode(&cfg); err != nil {
		panic(err)
	}

	registry, cancel, err := consul.RegisterService(cfg.APIConfig.Port, "localhost", serviceName, cfg.ConsulConfig.Host)
	if err != nil {
		log.Fatalf("failed to register service: %v", err)
	}
	defer cancel()
	defer registry.Deregister(context.Background(), discovery.GenerateInstanceID(serviceName), serviceName)

	repo, err := dbmysql.New()
	if err != nil {
		panic(err)
	}
	service := metadata.New(repo)
	h := grpcMetadata.New(service)

	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", cfg.APIConfig.Port))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()

	gen.RegisterMetadataServiceServer(server, h)
	server.Serve(listener)
}

func implRest() {
	/*var port int

	flag.IntVar(&port, "port", 6060, "Api handler port")
	flag.Parse()
	*/
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
