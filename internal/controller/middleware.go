package controller

import (
	"encoding/base64"
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
		decodedToken, err := base64.StdEncoding.DecodeString(authToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		isAuth, err := c.redis.Get(string(decodedToken)).Result()
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, err.Error())
			return
		}
		if isAuth == "" {
			ctx.Writer.Header().Add("Authorization", "")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		ctx.Next()
	}
}
