package services

import (
	"github.com/sagar-rathod-devops/do-host-network/internal/models"
	"github.com/sagar-rathod-devops/do-host-network/internal/repositories"
)

type AdminService interface {
	ModeratePost(postID, adminID string) error
	BanUser(userID, adminID string) error
}

type AnalyticsService interface {
	GetPostInteractions() ([]models.PostInteraction, error)
	GetUserAnalytics() ([]models.UserAnalytics, error)
}

type adminService struct {
	repo repositories.AdminRepository
}

type analyticsService struct {
	repo repositories.AnalyticsRepository
}

func NewAdminService(repo repositories.AdminRepository) AdminService {
	return &adminService{repo: repo}
}

func NewAnalyticsService(repo repositories.AnalyticsRepository) AnalyticsService {
	return &analyticsService{repo: repo}
}

func (s *adminService) ModeratePost(postID, adminID string) error {
	// Call repository to log the moderation action
	err := s.repo.LogAdminAction(adminID, postID, "post", "moderate")
	if err != nil {
		return err
	}

	// Implement logic for post moderation (e.g., mark the post as moderated, etc.)
	return nil
}

func (s *adminService) BanUser(userID, adminID string) error {
	// Call repository to log the ban action
	err := s.repo.LogAdminAction(adminID, userID, "user", "ban")
	if err != nil {
		return err
	}

	// Implement logic for banning the user (e.g., update their status to banned)
	return nil
}

func (s *analyticsService) GetPostInteractions() ([]models.PostInteraction, error) {
	return s.repo.GetPostInteractions()
}

func (s *analyticsService) GetUserAnalytics() ([]models.UserAnalytics, error) {
	return s.repo.GetUserAnalytics()
}
