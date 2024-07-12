package middleware

import (
	"gin-api/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// JWTAuthMiddleware creates a Gin middleware for JWT authentication.
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the JWT token from the header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  "Authorization header is required",
			})
			return
		}

		// The header format must be Bearer
		fields := strings.Fields(authHeader)
		if len(fields) < 2 || fields[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  "Authorization header format must be Bearer {token}",
			})
			return
		}

		tokenString := fields[1]
		tokenClaims, err := util.ParseToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  "Invalid token",
			})
			return
		}

		c.Set("userId", tokenClaims.UserId)

		c.Next()
	}
}
