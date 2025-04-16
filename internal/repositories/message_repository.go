package repositories

import (
	"database/sql"
	"time"

	"github.com/sagar-rathod-devops/do-host-network/internal/models"
)

type MessageRepository struct {
	DB *sql.DB
}

func NewMessageRepository(db *sql.DB) *MessageRepository {
	return &MessageRepository{DB: db}
}

func (r *MessageRepository) CreateMessage(message *models.Message) error {
	query := `
        INSERT INTO messages (sender_id, receiver_id, content, is_read, created_at)
        VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err := r.DB.QueryRow(query, message.SenderID, message.ReceiverID, message.Content, message.IsRead, time.Now()).
		Scan(&message.ID)
	return err
}

func (r *MessageRepository) GetConversation(senderID, receiverID string) ([]models.Message, error) {
	var messages []models.Message
	query := `
        SELECT id, sender_id, receiver_id, content, is_read, created_at
        FROM messages
        WHERE (sender_id = $1 AND receiver_id = $2)
           OR (sender_id = $2 AND receiver_id = $1)
        ORDER BY created_at ASC
    `
	rows, err := r.DB.Query(query, senderID, receiverID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var message models.Message
		if err := rows.Scan(&message.ID, &message.SenderID, &message.ReceiverID, &message.Content, &message.IsRead, &message.CreatedAt); err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}

	return messages, nil
}

func (r *MessageRepository) MarkMessagesAsRead(senderID, receiverID string) error {
	query := `
        UPDATE messages
        SET is_read = TRUE
        WHERE sender_id = $1 AND receiver_id = $2 AND is_read = FALSE
    `
	_, err := r.DB.Exec(query, senderID, receiverID)
	return err
}
