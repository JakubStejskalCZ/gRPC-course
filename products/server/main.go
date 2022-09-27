package main

import (
	"log"
	"net"

	pb "github.com/JakubStejskalCZ/gRPC-course/products/proto"
	"github.com/JakubStejskalCZ/gRPC-course/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	pb.ProductsServiceServer
	config util.Config
}

func main() {
	config, err := util.LoadConfig(".")

	if err != nil {
		log.Fatalf("Failed to load config: %v\n", err)
	}

	addr := config.GRPCServerAddress

	listener, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	defer listener.Close()

	log.Printf("Listening at %s\n", addr)

	opts := []grpc.ServerOption{}

	tls := config.Tls
	if tls {
		certFile := "ssl/server.crt"
		keyFile := "ssl/server.pem"
		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)

		if err != nil {
			log.Fatalf("Failed loading certificates: %v\n", err)
		}
		opts = append(opts, grpc.Creds(creds))
	}

	server := grpc.NewServer(opts...)
	pb.RegisterProductsServiceServer(server, &Server{
		config: config,
	})

	reflection.Register(server)

	if err = server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v\n", err)
	}
}
