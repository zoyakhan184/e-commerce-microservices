package handlers

import (
	"bff-service/clients"
	authpb "bff-service/proto/auth"
	"bff-service/utils"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// POST /api/auth/register
func Register(c *gin.Context) {
	var req authpb.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request")
		return
	}
	resp, err := clients.AuthClient().Register(c, &req)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(c, http.StatusCreated, resp)
}

// POST /api/auth/login
func Login(c *gin.Context) {
	var req authpb.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request")
		return
	}
	resp, err := clients.AuthClient().Login(c, &req)
	if err != nil {
		utils.RespondWithError(c, http.StatusUnauthorized, err.Error())
		return
	}
	utils.RespondWithJSON(c, http.StatusOK, resp)
}

// POST /api/auth/forgot-password
func ForgotPassword(c *gin.Context) {
	var req authpb.ForgotPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request")
		return
	}
	resp, err := clients.AuthClient().ForgotPassword(c, &req)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(c, http.StatusOK, resp)
}

// POST /api/auth/reset-password
func ResetPassword(c *gin.Context) {
	var req authpb.ResetPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request")
		return
	}
	resp, err := clients.AuthClient().ResetPassword(c, &req)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}
	utils.RespondWithJSON(c, http.StatusOK, resp)
}

// POST /api/auth/validate-token
func ValidateToken(c *gin.Context) {
	var token string

	// Try to read from JSON body
	var req authpb.ValidateTokenRequest
	if err := c.ShouldBindJSON(&req); err == nil && req.Token != "" {
		token = req.Token
	} else {
		// Try reading from Authorization header
		authHeader := c.GetHeader("Authorization")
		if strings.HasPrefix(authHeader, "Bearer ") {
			token = strings.TrimPrefix(authHeader, "Bearer ")
		}
	}

	if token == "" {
		utils.RespondWithError(c, http.StatusBadRequest, "Token is required")
		return
	}

	resp, err := clients.AuthClient().ValidateToken(c, &authpb.ValidateTokenRequest{Token: token})
	if err != nil {
		utils.RespondWithError(c, http.StatusUnauthorized, err.Error())
		return
	}

	utils.RespondWithJSON(c, http.StatusOK, resp)
}

func ChangePassword(c *gin.Context) {
	var jsonReq struct {
		CurrentPassword string `json:"currentPassword"`
		NewPassword     string `json:"newPassword"`
	}
	if err := c.ShouldBindJSON(&jsonReq); err != nil {
		log.Println("[ChangePassword] ❌ Invalid request:", err)
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request")
		return
	}

	userID, exists := c.Get("user_id")
	if !exists {
		log.Println("[ChangePassword] ❌ Unauthorized: user_id missing")
		utils.RespondWithError(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	log.Printf("[ChangePassword] 🔐 Password change requested for user: %s", userID)
	log.Printf("[ChangePassword] Current password length: %d", len(jsonReq.CurrentPassword))
	log.Printf("[ChangePassword] New password length: %d", len(jsonReq.NewPassword))

	// Construct gRPC request
	req := &authpb.ChangePasswordRequest{
		UserId:          userID.(string),
		CurrentPassword: jsonReq.CurrentPassword,
		NewPassword:     jsonReq.NewPassword,
	}

	resp, err := clients.AuthClient().ChangePassword(c, req)
	if err != nil {
		log.Printf("[ChangePassword] ❌ Error from auth-service: %v", err)
		utils.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	log.Printf("[ChangePassword] ✅ Password changed successfully for user: %s", req.UserId)
	utils.RespondWithJSON(c, http.StatusOK, resp)
}
