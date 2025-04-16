// services/user_profile_service.go
package services

import (
	"time"

	"github.com/sagar-rathod-devops/do-host-network/internal/models"
	"github.com/sagar-rathod-devops/do-host-network/internal/repositories"
)

type UserProfileService struct {
	Repo repositories.UserProfileRepository
}

func NewUserProfileService(repo repositories.UserProfileRepository) *UserProfileService {
	return &UserProfileService{Repo: repo}
}

func (s *UserProfileService) GetUserProfile(id string) (*models.UserProfile, string, error) {
	profile, email, err := s.Repo.GetByID(id)
	return profile, email, err
}

func (s *UserProfileService) CreateOrUpdateUserProfile(profile *models.UserProfile) (*models.UserProfile, error) {
	existingProfile, _, err := s.Repo.GetByID(profile.ID)
	if err != nil { // If profile doesn't exist, create it
		newProfile, err := s.Repo.Create(profile)
		if err != nil {
			return nil, err
		}
		return newProfile, nil
	}

	// Update existing profile
	profile.CreatedAt = existingProfile.CreatedAt
	profile.UpdatedAt = time.Now()
	return s.Repo.Create(profile)
}

func (s *UserProfileService) DeleteUserProfile(id string) error {
	return s.Repo.Delete(id)
}
