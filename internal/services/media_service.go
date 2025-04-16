package services

import (
	"github.com/sagar-rathod-devops/do-host-network/internal/models"
	"github.com/sagar-rathod-devops/do-host-network/internal/repositories"
)

type MediaService struct {
	Repo *repositories.MediaRepository
}

func NewMediaService(repo *repositories.MediaRepository) *MediaService {
	return &MediaService{Repo: repo}
}

func (s *MediaService) CreateMedia(media *models.MediaFile) error {
	return s.Repo.CreateMedia(media)
}

func (s *MediaService) GetMediaByUserID(userID string) ([]models.MediaFile, error) {
	return s.Repo.GetMediaByUserID(userID)
}
