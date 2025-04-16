package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sagar-rathod-devops/do-host-network/internal/models"
	"github.com/sagar-rathod-devops/do-host-network/internal/services"
)

type NotificationController struct {
	Service *services.NotificationService
}

func NewNotificationController(service *services.NotificationService) *NotificationController {
	return &NotificationController{Service: service}
}

func (c *NotificationController) GetNotifications(ctx *gin.Context) {
	userID := ctx.Param("userID")

	notifications, err := c.Service.GetNotifications(ctx, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(notifications) == 0 {
		ctx.JSON(http.StatusOK, []models.Notification{}) // Return empty array instead of null
		return
	}

	ctx.JSON(http.StatusOK, notifications)
}

func (c *NotificationController) MarkAsRead(ctx *gin.Context) {
	userID := ctx.Param("userID")

	count, err := c.Service.MarkAsRead(ctx, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if count == 0 {
		ctx.JSON(http.StatusOK, gin.H{"message": "No more unread notifications"})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"message": "Notifications marked as read"})
	}
}
