package services

import (
	"log"

	"github.com/sagar-rathod-devops/do-host-network/helpers"
	"github.com/sagar-rathod-devops/do-host-network/internal/models"
	"github.com/sagar-rathod-devops/do-host-network/internal/repositories"
)

type MessageService struct {
	repo *repositories.MessageRepository
}

func NewMessageService(repo *repositories.MessageRepository) *MessageService {
	return &MessageService{repo: repo}
}

func (s *MessageService) SendMessage(message *models.Message) error {
	// Save the message to the database
	if err := s.repo.CreateMessage(message); err != nil {
		return err
	}

	// Trigger the webhook
	if message.WebhookURL != "" {
		payload := helpers.WebhookPayload{
			ID:         message.ID,
			SenderID:   message.SenderID,
			ReceiverID: message.ReceiverID,
			Content:    message.Content,
			CreatedAt:  message.CreatedAt,
		}
		if err := helpers.TriggerWebhook(message.WebhookURL, payload); err != nil {
			log.Printf("Failed to trigger webhook: %v", err)
		}
	}

	return nil
}

func (s *MessageService) GetConversation(senderID, receiverID string) ([]models.Message, error) {
	// Mark messages as read
	if err := s.repo.MarkMessagesAsRead(senderID, receiverID); err != nil {
		return nil, err
	}

	// Fetch the conversation
	return s.repo.GetConversation(senderID, receiverID)
}
