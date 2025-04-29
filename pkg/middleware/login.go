package middleware

import (
	"context"
	"net/http"

	"github.com/g0shi4ek/test_project/cmd/clients"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(ctx context.Context, client *clients.AuthClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Autorization")
		if token == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "Authorization header required"})
			return
		}

		valid, err := client.TokenRepo.IsValidToken(token)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if !valid {
			//валидация токена
		}

		c.Next()
	}
}
