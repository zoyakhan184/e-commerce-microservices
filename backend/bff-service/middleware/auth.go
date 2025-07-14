package middleware

import (
	"bff-service/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var token string

		authHeader := c.GetHeader("Authorization")
		if strings.HasPrefix(authHeader, "Bearer ") {
			token = strings.TrimPrefix(authHeader, "Bearer ")
			println("[AuthMiddleware] Using Authorization header token")
		} else {
			cookieToken, err := c.Cookie("token")
			if err == nil && cookieToken != "" {
				token = cookieToken
				println("[AuthMiddleware] Using cookie token")
			} else {
				println("[AuthMiddleware] No token found")
			}
		}

		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
			c.Abort()
			return
		}

		uid, role, err := utils.ValidateJWT(token)
		if err != nil {
			println("[AuthMiddleware] Token validation failed:", err.Error())
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		println("[AuthMiddleware] Authenticated:", uid, role)

		c.Set("user_id", uid)
		c.Set("role", role)
		c.Next()
	}
}
