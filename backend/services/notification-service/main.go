package main

import (
	"log"
	"net"
	"notification-service/handlers"
	notificationpb "notification-service/proto"
	"notification-service/rabbitmq"
	"notification-service/repository"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	_ = godotenv.Load("../.env")

	repo := repository.NewNotificationRepo()
	rabbitmq.StartRabbitMQConsumer(repo)

	lis, _ := net.Listen("tcp", ":50056")
	grpcServer := grpc.NewServer()
	notificationpb.RegisterNotificationServiceServer(grpcServer, &handlers.NotificationService{Repo: repo})

	log.Println("📢 Notification Service running on port 50056")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("gRPC serve failed:", err)
	}
}
