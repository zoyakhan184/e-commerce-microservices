package main

import (
	"log"
	"net"
	"order-service/handlers"
	orderpb "order-service/proto"
	"order-service/rabbitmq"
	"order-service/repository"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	_ = godotenv.Load("../.env")

	repo := repository.NewOrderRepo()
	rabbitmq.InitRabbitMQ()
	rabbitmq.StartConsumer(repo)

	lis, _ := net.Listen("tcp", ":50057")
	grpcServer := grpc.NewServer()
	orderpb.RegisterOrderServiceServer(grpcServer, &handlers.OrderService{Repo: repo})

	log.Println("📦 Order Service running on port 50057")
	grpcServer.Serve(lis)
}
