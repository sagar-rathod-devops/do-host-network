package services

import (
	"github.com/sagar-rathod-devops/do-host-network/internal/models"
	"github.com/sagar-rathod-devops/do-host-network/internal/repositories"
)

type EducationService struct {
	Repo *repositories.EducationRepository
}

func (s *EducationService) GetEducationByID(id string) (*models.Education, error) {
	return s.Repo.GetEducationByID(id)
}

func (s *EducationService) CreateEducation(edu *models.Education) error {
	return s.Repo.CreateEducation(edu)
}

func (s *EducationService) DeleteEducationByID(id string) error {
	return s.Repo.DeleteEducationByID(id)
}

func (s *EducationService) ListEducationByUserID(userID string) ([]*models.Education, error) {
	return s.Repo.ListEducationByUserID(userID)
}
