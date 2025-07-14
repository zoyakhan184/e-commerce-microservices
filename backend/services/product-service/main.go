package main

import (
	"log"
	"net"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"

	"product-service/handlers"
	productpb "product-service/proto"
	"product-service/repository"
)

func main() {
	// Load env vars
	if err := godotenv.Load("../.env"); err != nil {
		log.Println("No .env file found, relying on system vars")
	}

	// Initialize repository
	repo := repository.NewPostgresProductRepo()
	handler := &handlers.ProductService{Repo: repo}

	// Start gRPC server
	lis, err := net.Listen("tcp", ":50059")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	productpb.RegisterProductServiceServer(grpcServer, handler)

	log.Println("Product Service is running on port :50059")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
