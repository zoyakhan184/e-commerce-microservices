package service

import (
	"context"
	"errors"
	"strconv"

	"github.com/zoyakhan1004/e-commerce-microservices/user-service/internal/models"
	"github.com/zoyakhan1004/e-commerce-microservices/user-service/internal/repository"
	"github.com/zoyakhan1004/e-commerce-microservices/user-service/proto"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserService struct {
	repo *repository.UserRepository
	proto.UnimplementedUserServiceServer
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Register(ctx context.Context, req *proto.RegisterRequest) (*proto.RegisterResponse, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to hash password: %v", err)
	}

	user := &models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	if err := s.repo.Create(user); err != nil {
		if errors.Is(err, repository.ErrEmailExists) {
			return nil, status.Errorf(codes.AlreadyExists, "email already exists")
		}
		return nil, status.Errorf(codes.Internal, "failed to create user: %v", err)
	}
	return &proto.RegisterResponse{UserId: strconv.FormatUint(uint64(user.ID), 10)}, nil
}

func (s *UserService) Login(ctx context.Context, req *proto.LoginRequest) (*proto.LoginResponse, error) {
	user, err := s.repo.GetByEmail(req.Email)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			return nil, status.Errorf(codes.NotFound, "invalid credentials")
		}
		return nil, status.Errorf(codes.Internal, "failed to get user: %v", err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid credentials")
	}

	// In production, generate a real JWT token here
	token := "generated-jwt-token-placeholder"

	return &proto.LoginResponse{Token: token}, nil
}
