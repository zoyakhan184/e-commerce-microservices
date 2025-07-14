package clients

import (
	userpb "bff-service/proto/user"
	"fmt"
	"log"
	"os"
	"sync"

	"google.golang.org/grpc"
)

var userOnce sync.Once
var userClient userpb.UserServiceClient

func UserClient() userpb.UserServiceClient {
	userOnce.Do(func() {
		addr := os.Getenv("USER_SERVICE_ADDR")
		conn, err := grpc.Dial(addr, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("gRPC dial error (User): %v", err)
		}
		userClient = userpb.NewUserServiceClient(conn)
		fmt.Println("✅ User client ready")
	})
	return userClient
}
