package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *Controller) BasicAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")

		if authHeader == "" || len(authHeader) < 6 || authHeader[:6] != "Basic " {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		authToken := authHeader[6:]

		isAuth, err := c.redis.Get(authToken).Result()
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		if isAuth == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		ctx.Next()
	}
}
