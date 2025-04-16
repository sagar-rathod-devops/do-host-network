package repositories

import (
	"context"
	"database/sql"
	"log"

	"github.com/sagar-rathod-devops/do-host-network/internal/models"
)

type NotificationRepository struct {
	DB *sql.DB
}

func NewNotificationRepository(db *sql.DB) *NotificationRepository {
	return &NotificationRepository{DB: db}
}

func (r *NotificationRepository) GetNotifications(ctx context.Context, userID string) ([]models.Notification, error) {
	rows, err := r.DB.QueryContext(ctx, `
        SELECT id, user_id, message, link, is_read, created_at
        FROM notifications
        WHERE user_id = $1 AND is_read = FALSE
        ORDER BY created_at DESC
    `, userID)
	if err != nil {
		log.Printf("Query error: %v", err)
		return nil, err
	}
	defer rows.Close()

	var notifications []models.Notification
	for rows.Next() {
		var n models.Notification
		var link sql.NullString
		if err := rows.Scan(&n.ID, &n.UserID, &n.Message, &link, &n.IsRead, &n.CreatedAt); err != nil {
			log.Printf("Row scan error: %v", err)
			return nil, err
		}
		if link.Valid {
			n.Link = &link.String
		}
		notifications = append(notifications, n)
	}

	if len(notifications) == 0 {
		log.Println("No unread notifications found for user:", userID)
	}
	return notifications, nil
}

func (r *NotificationRepository) MarkAsRead(ctx context.Context, userID string) (int64, error) {
	// Check if there are unread notifications
	var count int64
	err := r.DB.QueryRowContext(ctx, `
        SELECT COUNT(*)
        FROM notifications
        WHERE user_id = $1 AND is_read = FALSE
    `, userID).Scan(&count)
	if err != nil {
		return 0, err
	}

	if count == 0 {
		return 0, nil // No unread notifications
	}

	// Mark notifications as read
	result, err := r.DB.ExecContext(ctx, `
        UPDATE notifications
        SET is_read = TRUE
        WHERE user_id = $1 AND is_read = FALSE
    `, userID)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}
