package main

import (
	"log"
	"net"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"

	"inventory-service/handlers"
	inventorypb "inventory-service/proto"
	"inventory-service/rabbitmq"
	"inventory-service/repository"
)

func main() {
	_ = godotenv.Load("../.env")

	repo := repository.NewPostgresInventoryRepo()
	handler := &handlers.InventoryService{Repo: repo}

	rabbitmq.InitRabbitMQ(repo)

	lis, err := net.Listen("tcp", ":50055")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	inventorypb.RegisterInventoryServiceServer(grpcServer, handler)

	log.Println("Inventory Service running on port :50055")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
