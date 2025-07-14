package clients

import (
	paymentpb "bff-service/proto/payment"
	"fmt"
	"log"
	"os"
	"sync"

	"google.golang.org/grpc"
)

var paymentOnce sync.Once
var paymentClient paymentpb.PaymentServiceClient

func PaymentClient() paymentpb.PaymentServiceClient {
	paymentOnce.Do(func() {
		addr := os.Getenv("PAYMENT_SERVICE_ADDR")
		conn, err := grpc.Dial(addr, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("gRPC dial error (Payment): %v", err)
		}
		paymentClient = paymentpb.NewPaymentServiceClient(conn)
		fmt.Println("✅ Payment client ready")
	})
	return paymentClient
}
