package repositories

import (
	"database/sql"
	"errors"

	"github.com/sagar-rathod-devops/do-host-network/internal/models"
)

type ConnectionRepository interface {
	CreateConnection(connection *models.Connection) error
	UpdateConnectionStatus(userID, connectedUserID, status string) error
	DeleteConnection(userID, connectedUserID string) error
}

type connectionRepository struct {
	db *sql.DB
}

func NewConnectionRepository(db *sql.DB) ConnectionRepository {
	return &connectionRepository{db: db}
}

func (r *connectionRepository) CreateConnection(connection *models.Connection) error {
	query := `
		INSERT INTO connections (user_id, connected_user_id, status)
		VALUES ($1, $2, $3)
	`
	_, err := r.db.Exec(query, connection.UserID, connection.ConnectedUserID, connection.Status)
	return err
}

func (r *connectionRepository) UpdateConnectionStatus(userID, connectedUserID, status string) error {
	query := `
		UPDATE connections
		SET status = $1, updated_at = NOW()
		WHERE user_id = $2 AND connected_user_id = $3
	`
	result, err := r.db.Exec(query, status, userID, connectedUserID)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("connection not found")
	}
	return nil
}

func (r *connectionRepository) DeleteConnection(userID, connectedUserID string) error {
	query := `
		DELETE FROM connections
		WHERE user_id = $1 AND connected_user_id = $2
	`
	result, err := r.db.Exec(query, userID, connectedUserID)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("connection not found")
	}
	return nil
}
