package services

import (
	"github.com/sagar-rathod-devops/do-host-network/internal/models"
	"github.com/sagar-rathod-devops/do-host-network/internal/repositories"
)

type ExperienceService interface {
	CreateExperience(exp *models.Experience) error
	GetExperienceByID(id string) (*models.Experience, error)
	DeleteExperience(id string) error
	GetAllExperiencesByUserID(userID string) ([]models.Experience, error)
}

type experienceService struct {
	repo repositories.ExperienceRepository
}

func NewExperienceService(repo repositories.ExperienceRepository) ExperienceService {
	return &experienceService{repo: repo}
}

func (s *experienceService) CreateExperience(exp *models.Experience) error {
	return s.repo.CreateExperience(exp)
}

func (s *experienceService) GetExperienceByID(id string) (*models.Experience, error) {
	return s.repo.GetExperienceByID(id)
}

func (s *experienceService) DeleteExperience(id string) error {
	return s.repo.DeleteExperience(id)
}

func (s *experienceService) GetAllExperiencesByUserID(userID string) ([]models.Experience, error) {
	return s.repo.GetAllExperiencesByUserID(userID)
}
