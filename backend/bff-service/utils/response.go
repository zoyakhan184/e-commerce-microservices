package utils

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func RespondWithError(c *gin.Context, code int, msg string) {
	c.JSON(code, gin.H{"error": msg})
}

func RespondWithJSON(c *gin.Context, code int, data interface{}) {
	c.JSON(code, data)
}

func ValidateJWT(tokenStr string) (string, string, error) {
	// Parse token without verifying signature
	parser := jwt.Parser{
		SkipClaimsValidation: true, // skip expiry & other registered claims
	}

	claims := jwt.MapClaims{}
	_, _, err := parser.ParseUnverified(tokenStr, claims)
	if err != nil {
		return "", "", fmt.Errorf("failed to parse token: %v", err)
	}

	// Extract user_id and role from claims
	userID, ok1 := claims["user_id"].(string)
	role, ok2 := claims["role"].(string)
	if !ok1 || !ok2 {
		return "", "", fmt.Errorf("missing or invalid user_id/role in token claims")
	}

	// Optional: log warning that signature was NOT verified
	fmt.Println("⚠️  WARNING: JWT signature was not verified")

	return userID, role, nil
}

// ✅ Secure, safe and configurable JWT validator
// func ValidateJWT(tokenStr string) (string, string, error) {
// 	secret := os.Getenv("JWT_SECRET")
// 	if secret == "" {
// 		return "", "", fmt.Errorf("JWT_SECRET not set")
// 	}

// 	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
// 		// Validate signing method
// 		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("unexpected signing method")
// 		}
// 		return []byte(secret), nil
// 	})

// 	if err != nil || !token.Valid {
// 		return "", "", fmt.Errorf("invalid token: %v", err)
// 	}

// 	claims, ok := token.Claims.(jwt.MapClaims)
// 	if !ok {
// 		return "", "", fmt.Errorf("invalid token claims")
// 	}

// 	// Expiry check
// 	if exp, ok := claims["exp"].(float64); ok {
// 		if int64(exp) < time.Now().Unix() {
// 			return "", "", fmt.Errorf("token expired")
// 		}
// 	}

// 	// Safely extract user_id and role
// 	userID, ok1 := claims["user_id"].(string)
// 	role, ok2 := claims["role"].(string)
// 	if !ok1 || !ok2 {
// 		return "", "", fmt.Errorf("missing claims: user_id or role")
// 	}

// 	return userID, role, nil
// }
