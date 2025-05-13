package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RoleAuthMiddleware(role string) gin.HandlerFunc{
	return func(c *gin.Context) {
		user, exists := c.Get("user")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
			c.Abort()
			return
		}

		claims, ok := user.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid user data"})
			c.Abort()
			return
		}

		if claims["role"] != role {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "access forbidden"})
			c.Abort()
			return
		}

		c.Next()
	}
}