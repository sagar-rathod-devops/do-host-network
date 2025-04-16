package repositories

import (
	"database/sql"
	"errors"

	"github.com/sagar-rathod-devops/do-host-network/internal/models"
)

type GroupRepository struct {
	DB *sql.DB
}

func NewGroupRepository(db *sql.DB) *GroupRepository {
	return &GroupRepository{DB: db}
}

func (repo *GroupRepository) CreateGroup(group *models.Group) error {
	query := `
		INSERT INTO groups (name, description, group_picture_url) 
		VALUES ($1, $2, $3) RETURNING id, created_at, updated_at`
	err := repo.DB.QueryRow(query, group.Name, group.Description, group.GroupPictureURL).
		Scan(&group.ID, &group.CreatedAt, &group.UpdatedAt)
	return err
}

func (repo *GroupRepository) GetGroupByID(id string) (*models.Group, error) {
	query := `SELECT id, name, description, group_picture_url, created_at, updated_at FROM groups WHERE id = $1`
	row := repo.DB.QueryRow(query, id)

	group := &models.Group{}
	err := row.Scan(&group.ID, &group.Name, &group.Description, &group.GroupPictureURL, &group.CreatedAt, &group.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, errors.New("group not found")
	}
	return group, err
}

func (repo *GroupRepository) AddGroupMember(member *models.GroupMember) error {
	query := `
		INSERT INTO group_members (group_id, user_id, role) 
		VALUES ($1, $2, $3) RETURNING id, joined_at`
	err := repo.DB.QueryRow(query, member.GroupID, member.UserID, member.Role).
		Scan(&member.ID, &member.JoinedAt)
	return err
}
