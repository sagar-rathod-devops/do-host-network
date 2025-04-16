package services

import (
	"github.com/sagar-rathod-devops/do-host-network/internal/models"
	"github.com/sagar-rathod-devops/do-host-network/internal/repositories"
)

type GroupService struct {
	Repo *repositories.GroupRepository
}

func NewGroupService(repo *repositories.GroupRepository) *GroupService {
	return &GroupService{Repo: repo}
}

func (svc *GroupService) CreateGroup(group *models.Group) error {
	return svc.Repo.CreateGroup(group)
}

func (svc *GroupService) GetGroupByID(id string) (*models.Group, error) {
	return svc.Repo.GetGroupByID(id)
}

func (svc *GroupService) AddGroupMember(member *models.GroupMember) error {
	return svc.Repo.AddGroupMember(member)
}
