package clients

import (
	adminpb "bff-service/proto/admin"
	"fmt"
	"log"
	"os"
	"sync"

	"google.golang.org/grpc"
)

var adminOnce sync.Once
var adminClient adminpb.AdminServiceClient

func AdminClient() adminpb.AdminServiceClient {
	adminOnce.Do(func() {
		addr := os.Getenv("ADMIN_SERVICE_ADDR")
		conn, err := grpc.Dial(addr, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("gRPC dial error (Admin): %v", err)
		}
		adminClient = adminpb.NewAdminServiceClient(conn)
		fmt.Println("✅ Admin client ready")
	})
	return adminClient
}
