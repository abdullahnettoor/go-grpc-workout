package order

import (
	"github.com/abdullahnettoor/go-grpc-workout/tree/main/api-gateway/pkg/auth"
	"github.com/abdullahnettoor/go-grpc-workout/tree/main/api-gateway/pkg/config"
	"github.com/abdullahnettoor/go-grpc-workout/tree/main/api-gateway/pkg/order/routes"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	router := r.Group("/order").Use(a.AuthRequired)
	router.POST("/", svc.CreateOrder)
}

func (svc *ServiceClient) CreateOrder(ctx *gin.Context) {
	routes.CreateOrder(ctx, svc.Client)
}
