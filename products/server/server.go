package main

import (
	"context"
	"log"

	pb "github.com/JakubStejskalCZ/gRPC-course/products/proto"
)

func (server *Server) GetProducts(ctx context.Context, request *pb.GetProductsRequest) (*pb.GetProductsResponse, error) {
	log.Printf("Called: GetProducts() with: %v\n", request)

	return &pb.GetProductsResponse{}, nil
}

func (server *Server) GetProduct(ctx context.Context, request *pb.GetProductRequest) (*pb.GetProductResponse, error) {
	log.Printf("Called: GetProduct() with: %v\n", request)

	return &pb.GetProductResponse{}, nil
}

func (server *Server) CreateProduct(ctx context.Context, request *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	log.Printf("Called: CreateProduct() with: %v\n", request)

	return &pb.CreateProductResponse{}, nil
}

func (server *Server) DeleteProduct(ctx context.Context, request *pb.DeleteProductRequest) (*pb.DeleteProductResponse, error) {
	log.Printf("Called: DeleteProduct() with: %v\n", request)

	return &pb.DeleteProductResponse{}, nil
}

func (server *Server) UpdateProduct(ctx context.Context, request *pb.UpdateProductRequest) (*pb.UpdateProductResponse, error) {
	log.Printf("Called: UpdateProduct() with: %v\n", request)

	return &pb.UpdateProductResponse{}, nil
}
