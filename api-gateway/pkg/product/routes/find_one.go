package routes

import (
	"context"
	"net/http"
	"strconv"

	"github.com/abdullahnettoor/go-grpc-workout/tree/main/api-gateway/pkg/product/pb"
	"github.com/gin-gonic/gin"
)

func FineOne(ctx *gin.Context, c pb.ProductServiceClient) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 32)

	res, err := c.FindOne(context.Background(), &pb.FindOneRequest{
		ProductId: int64(id),
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(int(res.Status), &res)
}
