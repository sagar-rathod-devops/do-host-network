package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sagar-rathod-devops/do-host-network/internal/models"
	"github.com/sagar-rathod-devops/do-host-network/internal/services"
)

type MessageController struct {
	service *services.MessageService
}

func NewMessageController(service *services.MessageService) *MessageController {
	return &MessageController{service: service}
}

func (c *MessageController) SendMessage(ctx *gin.Context) {
	var message models.Message
	if err := ctx.ShouldBindJSON(&message); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.service.SendMessage(&message); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Message sent successfully"})
}

func (c *MessageController) GetConversation(ctx *gin.Context) {
	receiverID := ctx.Param("id")      // Extract the receiver_id from the URL path
	senderID := ctx.Query("sender_id") // Extract the sender_id from the query parameter

	if receiverID == "" || senderID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "receiver_id and sender_id are required"})
		return
	}

	// Fetch the conversation
	messages, err := c.service.GetConversation(senderID, receiverID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch messages"})
		return
	}

	ctx.JSON(http.StatusOK, messages)
}
