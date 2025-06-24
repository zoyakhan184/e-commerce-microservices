package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zoyakhan1004/e-commerce-microservices/bff/internal/clients"
	"github.com/zoyakhan1004/e-commerce-microservices/bff/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserHandler struct {
	grpcClients *clients.GrpcClients
}

func NewUserHandler(clients *clients.GrpcClients) *UserHandler {
	return &UserHandler{grpcClients: clients}
}

func (h *UserHandler) Register(c *gin.Context) {
	var req struct {
		Name     string `json:"name" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=8"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.grpcClients.UserClient.Register(c.Request.Context(), &proto.RegisterRequest{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	})

	if err != nil {
		handleGrpcError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"user_id": res.UserId,
	})
}

func (h *UserHandler) Login(c *gin.Context) {
	var req struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.grpcClients.UserClient.Login(c.Request.Context(), &proto.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
	})

	if err != nil {
		handleGrpcError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": res.Token,
	})
}

func handleGrpcError(c *gin.Context, err error) {
	st, ok := status.FromError(err)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	switch st.Code() {
	case codes.AlreadyExists:
		c.JSON(http.StatusConflict, gin.H{"error": st.Message()})
	case codes.NotFound, codes.Unauthenticated:
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
	}
}
