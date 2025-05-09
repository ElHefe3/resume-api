package middleware

import (
	"net/http"

	"github.com/ElHefe3/resume-api/internal/config"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authToken := config.Cfg.AuthToken

		token := c.GetHeader("Authorization")
		if token == "" || token != authToken {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
			})
			return
		}

		// continue to handler
		c.Next()
	}
}
