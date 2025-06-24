package clients

import (
	"time"

	"github.com/zoyakhan1004/e-commerce-microservices/bff/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GrpcClients struct {
	UserClient proto.UserServiceClient
}

func NewGrpcClients() (*GrpcClients, error) {
	// User Service Connection
	userConn, err := grpc.Dial(
		"user-service:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithTimeout(5*time.Second),
	)
	if err != nil {
		return nil, err
	}

	return &GrpcClients{
		UserClient: proto.NewUserServiceClient(userConn),
	}, nil
}
