package auth

import (
	"fmt"

	"github.com/abdullahnettoor/go-grpc-workout/tree/main/api-gateway/pkg/auth/pb"
	"github.com/abdullahnettoor/go-grpc-workout/tree/main/api-gateway/pkg/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Client pb.AuthServiceClient
}

func InitServeceClient(c *config.Config) pb.AuthServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.AuthSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewAuthServiceClient(cc)
}
