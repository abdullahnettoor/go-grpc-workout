package routes

import (
	"context"
	"net/http"

	"github.com/abdullahnettoor/go-grpc-workout/tree/main/api-gateway/pkg/order/pb"
	"github.com/gin-gonic/gin"
)

type CreateOrderRequestBody struct {
	ProductId int64 `json:"productid"`
	Quantity  int64 `json:"quantity"`
}

func CreateOrder(ctx *gin.Context, c pb.OrderServiceClient) {
	body := CreateOrderRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	userId, _ := ctx.Get("userId")

	res, err := c.CreateOrder(context.Background(), &pb.CreateOrderRequest{
		ProductId: body.ProductId,
		Quantity:  body.Quantity,
		UserId:    userId.(int64),
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(201, &res)
}
