package services

import (
	"context"

	"github.com/sagar-rathod-devops/do-host-network/internal/models"
	"github.com/sagar-rathod-devops/do-host-network/internal/repositories"
)

type NotificationService struct {
	Repo *repositories.NotificationRepository
}

func NewNotificationService(repo *repositories.NotificationRepository) *NotificationService {
	return &NotificationService{Repo: repo}
}

func (s *NotificationService) GetNotifications(ctx context.Context, userID string) ([]models.Notification, error) {
	return s.Repo.GetNotifications(ctx, userID)
}

func (s *NotificationService) MarkAsRead(ctx context.Context, userID string) (int64, error) {
	return s.Repo.MarkAsRead(ctx, userID)
}
