package product

import (
	"fmt"

	"github.com/abdullahnettoor/go-grpc-workout/tree/main/api-gateway/pkg/config"
	"github.com/abdullahnettoor/go-grpc-workout/tree/main/api-gateway/pkg/product/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Client pb.ProductServiceClient
}

func InitServiceClient(c *config.Config) pb.ProductServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.ProductSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewProductServiceClient(cc)
}
