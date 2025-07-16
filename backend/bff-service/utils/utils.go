package utils

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

// RespondWithError sends a JSON error response
func RespondWithError(c *gin.Context, code int, msg string) {
	fmt.Printf("[RespondWithError] %d: %s\n", code, msg)
	c.JSON(code, gin.H{"error": msg})
}

// RespondWithJSON sends JSON response
func RespondWithJSON(c *gin.Context, code int, data interface{}) {
	fmt.Printf("[RespondWithJSON] %d: %+v\n", code, data)
	c.JSON(code, data)
}

// ParseInt safely converts string to int
func ParseInt(s string, defaultVal int) int {
	if val, err := strconv.Atoi(s); err == nil {
		return val
	}
	fmt.Printf("[ParseInt] Could not parse '%s', using default: %d\n", s, defaultVal)
	return defaultVal
}

// Dynamically get JWT_SECRET at runtime
func getJWTSecret() []byte {
	secret := os.Getenv("JWT_SECRET")
	fmt.Println("JWT_SECRET:", secret)
	if secret == "" {
		panic("JWT_SECRET is not set in environment")
	}
	return []byte(secret)
}

// ValidateJWT parses and verifies the JWT and returns userID and role
func ValidateJWT(tokenStr string) (string, string, error) {
	fmt.Println("[ValidateJWT] Starting token validation...")
	fmt.Println("[ValidateJWT] Token string:", tokenStr)

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		fmt.Printf("[ValidateJWT] Token Header: %+v\n", token.Header)
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return getJWTSecret(), nil
	})

	if err != nil {
		fmt.Println("[ValidateJWT] ❌ Error parsing token:", err)
		return "", "", fmt.Errorf("invalid token: %v", err)
	}

	if !token.Valid {
		fmt.Println("[ValidateJWT] ❌ Invalid token signature")
		return "", "", errors.New("token signature is invalid")
	}

	if claims.ExpiresAt == nil || time.Now().After(claims.ExpiresAt.Time) {
		fmt.Println("[ValidateJWT] ❌ Token is expired")
		return "", "", errors.New("token expired")
	}

	fmt.Printf("[ValidateJWT] ✅ Token valid. user_id: %s, role: %s\n", claims.UserID, claims.Role)
	return claims.UserID, claims.Role, nil
}
