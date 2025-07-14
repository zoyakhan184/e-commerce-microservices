package main

import (
    "cart-service/handlers"
    cartpb "cart-service/proto"
    "cart-service/repository"
    "github.com/joho/godotenv"
    "google.golang.org/grpc"
    "log"
    "net"
)

func main() {
    _ = godotenv.Load("../.env")

    repo := repository.NewPostgresCartRepo()
    handler := &handlers.CartService{Repo: repo}

    lis, err := net.Listen("tcp", ":50053")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    grpcServer := grpc.NewServer()
    cartpb.RegisterCartServiceServer(grpcServer, handler)

    log.Println("Cart Service is running on port :50053")
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
