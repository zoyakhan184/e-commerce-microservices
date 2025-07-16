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

		// ✅ Step 1: Extract Token
		authHeader := c.GetHeader("Authorization")
		if strings.HasPrefix(authHeader, "Bearer ") {
			token = strings.TrimPrefix(authHeader, "Bearer ")
			println("[AuthMiddleware] ✅ Using token from Authorization header:", token)
		} else {
			cookieToken, err := c.Cookie("token")
			if err == nil && cookieToken != "" {
				token = cookieToken
				println("[AuthMiddleware] ✅ Using token from cookie:", token)
			} else {
				println("[AuthMiddleware] ❌ No token found in Authorization header or cookie")
			}
		}

		// ✅ Step 2: Check for missing token
		if token == "" {
			println("[AuthMiddleware] ⛔ Token is missing in the request")
			utils.RespondWithError(c, http.StatusUnauthorized, "Missing token")
			c.Abort()
			return
		}

		// ✅ Step 3: Validate JWT
		println("[AuthMiddleware] 🔍 Validating token:", token)
		userID, role, err := utils.ValidateJWT(token)
		println("[AuthMiddleware] 🧪 Validation result - userID:", userID, ", role:", role, ", error:", err)

		if err != nil {
			println("[AuthMiddleware] ❌ Token validation failed:", err.Error())
			utils.RespondWithError(c, http.StatusUnauthorized, "Invalid token")
			c.Abort()
			return
		}

		// ✅ Step 4: Store user info in context
		println("[AuthMiddleware] ✅ Token valid. Authenticated user ID:", userID, ", role:", role)
		c.Set("user_id", userID)
		c.Set("role", role)

		c.Next()
	}
}
