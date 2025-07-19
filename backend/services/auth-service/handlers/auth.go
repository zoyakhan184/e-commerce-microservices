package handlers

import (
	"auth-service/models"
	authpb "auth-service/proto"
	"auth-service/rabbitmq"
	"auth-service/utils"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuthService struct {
	DB *gorm.DB
	authpb.UnimplementedAuthServiceServer
}

func (s *AuthService) Register(ctx context.Context, req *authpb.RegisterRequest) (*authpb.AuthResponse, error) {
	hashed, _ := utils.HashPassword(req.Password)
	user := models.User{
		ID:        uuid.New().String(),
		Name:      req.Name,
		Email:     req.Email,
		Password:  hashed,
		Role:      "user", // ✅ Force role to "user"
		CreatedAt: time.Now(),
	}

	if err := s.DB.Create(&user).Error; err != nil {
		return nil, err
	}

	// ✅ Emit structured user.registered event for admin activity
	rabbitmq.Publish("user.registered", map[string]interface{}{
		"user_id":   user.ID,
		"name":      user.Name,
		"email":     user.Email,
		"timestamp": user.CreatedAt.Format(time.RFC3339),
	})

	token, _ := utils.GenerateJWT(user.ID, user.Role)
	fmt.Println("generated token:", token)
	return &authpb.AuthResponse{Token: token, UserId: user.ID, Role: user.Role}, nil
}

func (s *AuthService) Login(ctx context.Context, req *authpb.LoginRequest) (*authpb.AuthResponse, error) {
	var user models.User
	fmt.Print("Login attempt for email: ", req.Email, "\n")
	if err := s.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		return nil, err
	}
	fmt.Println("User found:", user.Email)
	if !utils.CheckPasswordHash(req.Password, user.Password) {
		temphash, _ := utils.HashPassword(req.Password)
		fmt.Println("Password hash for user:", user.Email, "is", temphash)
		fmt.Println("Password mismatch for user:", user.Email)
		return nil, fmt.Errorf("invalid credentials")
	}
	fmt.Println("Password matched for user:", user.Email)
	token, _ := utils.GenerateJWT(user.ID, user.Role)
	return &authpb.AuthResponse{Token: token, UserId: user.ID, Role: user.Role}, nil
}

func (s *AuthService) ForgotPassword(ctx context.Context, req *authpb.ForgotPasswordRequest) (*authpb.GenericResponse, error) {
	var user models.User
	if err := s.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		return nil, err
	}

	resetToken := uuid.New().String()
	reset := models.PasswordReset{
		ID:        uuid.New().String(),
		UserID:    user.ID,
		Token:     resetToken,
		ExpiresAt: time.Now().Add(15 * time.Minute),
		Used:      false,
	}
	s.DB.Create(&reset)

	rabbitmq.Publish("auth.forgot_password", map[string]interface{}{
		"email": user.Email,
		"token": resetToken,
	})

	return &authpb.GenericResponse{Message: "Reset link sent"}, nil
}

func (s *AuthService) ResetPassword(ctx context.Context, req *authpb.ResetPasswordRequest) (*authpb.GenericResponse, error) {
	var reset models.PasswordReset
	if err := s.DB.Where("token = ?", req.Token).First(&reset).Error; err != nil {
		return nil, err
	}
	if reset.Used || reset.ExpiresAt.Before(time.Now()) {
		return nil, fmt.Errorf("invalid or expired token")
	}

	hashed, _ := utils.HashPassword(req.NewPassword)
	s.DB.Model(&models.User{}).Where("id = ?", reset.UserID).Update("password", hashed)
	s.DB.Model(&reset).Update("used", true)

	rabbitmq.Publish("auth.password_reset", map[string]interface{}{
		"user_id": reset.UserID,
	})

	return &authpb.GenericResponse{Message: "Password updated"}, nil
}

func (s *AuthService) ValidateToken(ctx context.Context, req *authpb.ValidateTokenRequest) (*authpb.ValidateTokenResponse, error) {
	fmt.Println("Validating token:", req.Token)
	claims, err := utils.ParseJWT(req.Token)
	if err != nil {
		return nil, fmt.Errorf("invalid token: %v", err)
	}

	return &authpb.ValidateTokenResponse{
		UserId: claims.UserID,
		Role:   claims.Role,
	}, nil
}

func (s *AuthService) ChangePassword(ctx context.Context, req *authpb.ChangePasswordRequest) (*authpb.ChangePasswordResponse, error) {
	// 1. Fetch user by ID
	var user models.User
	if err := s.DB.Where("id = ?", req.UserId).First(&user).Error; err != nil {
		return nil, fmt.Errorf("user not found")
	}

	// 2. Check if current password matches
	if !utils.CheckPasswordHash(req.CurrentPassword, user.Password) {
		return nil, fmt.Errorf("current password is incorrect")
	}

	// 3. Hash new password
	newHashed, err := utils.HashPassword(req.NewPassword)
	if err != nil {
		return nil, fmt.Errorf("failed to hash new password")
	}

	// 4. Update password in DB
	if err := s.DB.Model(&models.User{}).Where("id = ?", user.ID).Update("password", newHashed).Error; err != nil {
		return nil, fmt.Errorf("failed to update password")
	}

	// 5. Emit event (optional)
	rabbitmq.Publish("auth.password_changed", map[string]interface{}{
		"user_id": user.ID,
		"email":   user.Email,
	})

	// 6. Respond success
	return &authpb.ChangePasswordResponse{
		Message: "Password updated successfully",
	}, nil
}
