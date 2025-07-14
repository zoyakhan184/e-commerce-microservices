package main

import (
	"image-service/handlers"
	"image-service/proto"
	"image-service/repository"
	"log"
	"net"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	_ = godotenv.Load("../.env")
	repo := repository.NewImageRepo()
	grpcServer := grpc.NewServer()
	proto.RegisterImageServiceServer(grpcServer, &handlers.ImageService{Repo: repo})

	lis, err := net.Listen("tcp", ":50054")
	if err != nil {
		log.Fatal("Failed to listen:", err)
	}

	log.Println("📷 Image Service running on :50054")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("Failed to serve:", err)
	}
}
