package clients

import (
	authpb "bff-service/proto/auth"
	"fmt"
	"log"
	"os"
	"sync"

	"google.golang.org/grpc"
)

var authOnce sync.Once
var authClient authpb.AuthServiceClient

func AuthClient() authpb.AuthServiceClient {
	authOnce.Do(func() {
		addr := os.Getenv("AUTH_SERVICE_ADDR")
		conn, err := grpc.Dial(addr, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("gRPC dial error (Auth): %v", err)
		}
		authClient = authpb.NewAuthServiceClient(conn)
		fmt.Println("✅ Auth client ready")
	})
	return authClient
}
