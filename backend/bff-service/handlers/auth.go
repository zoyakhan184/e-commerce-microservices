package handlers

import (
	"bff-service/clients"
	authpb "bff-service/proto/auth"
	"bff-service/utils"
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
