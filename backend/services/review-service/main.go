package main

import (
	"log"
	"net"
	"review-service/handlers"
	reviewpb "review-service/proto"
	"review-service/repository"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	_ = godotenv.Load("../.env")

	repo := repository.NewReviewRepo()
	handler := &handlers.ReviewService{Repo: repo}

	lis, err := net.Listen("tcp", ":50060")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	reviewpb.RegisterReviewServiceServer(grpcServer, handler)

	log.Println("📝 Review Service is running on port :50060")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("gRPC failed to serve: %v", err)
	}
}
