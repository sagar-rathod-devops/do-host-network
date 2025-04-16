package services

import (
	"github.com/sagar-rathod-devops/do-host-network/internal/models"
	"github.com/sagar-rathod-devops/do-host-network/internal/repositories"
)

type ConnectionService interface {
	FollowUser(userID, connectedUserID string) error
	UnfollowUser(userID, connectedUserID string) error
	SendFriendRequest(userID, connectedUserID string) error
}

type connectionService struct {
	repo repositories.ConnectionRepository
}

func NewConnectionService(repo repositories.ConnectionRepository) ConnectionService {
	return &connectionService{repo: repo}
}

func (s *connectionService) FollowUser(userID, connectedUserID string) error {
	connection := &models.Connection{
		UserID:          userID,
		ConnectedUserID: connectedUserID,
		Status:          "follow",
	}
	return s.repo.CreateConnection(connection)
}

func (s *connectionService) UnfollowUser(userID, connectedUserID string) error {
	return s.repo.DeleteConnection(userID, connectedUserID)
}

func (s *connectionService) SendFriendRequest(userID, connectedUserID string) error {
	connection := &models.Connection{
		UserID:          userID,
		ConnectedUserID: connectedUserID,
		Status:          "pending",
	}
	return s.repo.CreateConnection(connection)
}
