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

	opts := []grpc.ServerOption{}

	tls := config.Tls

	log.Printf("TLS: %t\n", tls)

	if tls {
		certFile := "ssl/server.crt"
		keyFile := "ssl/server.pem"
		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)

		if err != nil {
			log.Fatalf("Failed loading certificates: %v\n", err)
		}

		opts = append(opts, grpc.Creds(creds))
	}

	log.Println("Starting gRPC server...")

	server := grpc.NewServer(opts...)

	log.Println("Registering ProductsServiceServer...")

	pb.RegisterProductsServiceServer(server, &Server{})

	log.Println("Reflection enabled...")

	reflection.Register(server)

	log.Printf("Listening at %s\n", addr)

	if err = server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
