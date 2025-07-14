package handlers

import (
	"context"
	"notification-service/config"
	"notification-service/models"
	notificationpb "notification-service/proto"
	"notification-service/repository"
	"time"

	"github.com/google/uuid"
)

type NotificationService struct {
	notificationpb.UnimplementedNotificationServiceServer
	Repo *repository.NotificationRepo
}

func (s *NotificationService) SendEmail(ctx context.Context, req *notificationpb.EmailRequest) (*notificationpb.GenericResponse, error) {
	err := config.SendEmail(req.To, req.Subject, req.Body)
	if err == nil {
		s.Repo.Save(&models.Notification{
			ID:        uuid.New().String(),
			UserID:    req.UserId,
			Title:     req.Subject,
			Message:   req.Body,
			Read:      false,
			CreatedAt: time.Now(),
		})
	}
	return &notificationpb.GenericResponse{Status: "sent"}, err
}

func (s *NotificationService) SendBulkEmail(ctx context.Context, req *notificationpb.BulkEmailRequest) (*notificationpb.GenericResponse, error) {
	config.SendBulkEmail(req.ToList, req.Subject, req.Body)
	return &notificationpb.GenericResponse{Status: "sent"}, nil
}
