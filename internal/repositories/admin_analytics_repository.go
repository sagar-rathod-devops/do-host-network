package repositories

import (
	"database/sql"
	"fmt"

	"github.com/sagar-rathod-devops/do-host-network/internal/models"
)

type AdminRepository interface {
	LogAdminAction(adminID, targetID, targetType, action string) error
}

type AnalyticsRepository interface {
	GetPostInteractions() ([]models.PostInteraction, error)
	GetUserAnalytics() ([]models.UserAnalytics, error)
}

type adminRepository struct {
	db *sql.DB
}

type analyticsRepository struct {
	db *sql.DB
}

func NewAdminRepository(db *sql.DB) AdminRepository {
	return &adminRepository{db: db}
}

func NewAnalyticsRepository(db *sql.DB) AnalyticsRepository {
	return &analyticsRepository{db: db}
}

func (r *analyticsRepository) GetPostInteractions() ([]models.PostInteraction, error) {
	query := `SELECT id, post_id, likes, comments, created_at FROM post_interactions`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var interactions []models.PostInteraction
	for rows.Next() {
		var interaction models.PostInteraction
		if err := rows.Scan(&interaction.ID, &interaction.PostID, &interaction.Likes, &interaction.Comments, &interaction.CreatedAt); err != nil {
			return nil, err
		}
		interactions = append(interactions, interaction)
	}
	return interactions, nil
}

func (r *analyticsRepository) GetUserAnalytics() ([]models.UserAnalytics, error) {
	query := `SELECT id, user_id, engagement_score, active, created_at FROM user_analytics`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var analytics []models.UserAnalytics
	for rows.Next() {
		var analytic models.UserAnalytics
		if err := rows.Scan(&analytic.ID, &analytic.UserID, &analytic.EngagementScore, &analytic.Active, &analytic.CreatedAt); err != nil {
			return nil, err
		}
		analytics = append(analytics, analytic)
	}
	return analytics, nil
}

func (r *adminRepository) LogAdminAction(adminID, targetID, targetType, action string) error {
	query := `
		INSERT INTO admin_logs (admin_id, target_id, target_type, action, created_at)
		VALUES ($1, $2, $3, $4, NOW())`
	_, err := r.db.Exec(query, adminID, targetID, targetType, action)
	if err != nil {
		return fmt.Errorf("failed to log admin action: %w", err)
	}
	return nil
}
