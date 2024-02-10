package order

import (
	"fmt"

	"github.com/abdullahnettoor/go-grpc-workout/tree/main/api-gateway/pkg/config"
	"github.com/abdullahnettoor/go-grpc-workout/tree/main/api-gateway/pkg/order/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Client pb.OrderServiceClient
}

func InitServiceClient(c *config.Config) pb.OrderServiceClient {

	cc, err := grpc.Dial(c.OrderSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Couldn't Connect:", err)
	}

	return pb.NewOrderServiceClient(cc)
}
