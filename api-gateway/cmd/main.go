package main

import (
	"log"

	"github.com/abdullahnettoor/go-grpc-workout/tree/main/api-gateway/pkg/auth"
	"github.com/abdullahnettoor/go-grpc-workout/tree/main/api-gateway/pkg/config"
	"github.com/abdullahnettoor/go-grpc-workout/tree/main/api-gateway/pkg/order"
	"github.com/abdullahnettoor/go-grpc-workout/tree/main/api-gateway/pkg/product"
	"github.com/gin-gonic/gin"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at Config:", err)
	}

	r := gin.Default()

	authSvc := *auth.RegisterRoutes(r, c)
	product.RegisterRoutes(r, c, &authSvc)
	order.RegisterRoutes(r, c, &authSvc)

	r.Run(c.Port)
}
