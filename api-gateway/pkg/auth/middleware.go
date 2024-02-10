package auth

import (
	"context"
	"net/http"
	"strings"

	"github.com/abdullahnettoor/go-grpc-workout/tree/main/api-gateway/pkg/auth/pb"
	"github.com/gin-gonic/gin"
)

type AuthMiddlewareConfig struct {
	svc *ServiceClient
}

func InitAuthMiddleware(svc *ServiceClient) AuthMiddlewareConfig {
	return AuthMiddlewareConfig{svc}
}

func (c *AuthMiddlewareConfig) AuthRequired(ctx *gin.Context) {
	token := ctx.Request.Header.Get("authorization")
	token = strings.TrimPrefix(token, "Bearer ")
	if token == "" {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	res, err := c.svc.Client.Validate(context.Background(), &pb.ValidateRequest{Token: token})
	if err != nil || res.Status != 200 {
		ctx.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	ctx.Set("userId", res.UserId)
	ctx.Next()
}
