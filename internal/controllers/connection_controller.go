// controllers/connection_controller.go
package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sagar-rathod-devops/do-host-network/internal/services"
)

type ConnectionController struct {
	service services.ConnectionService
}

func NewConnectionController(service services.ConnectionService) *ConnectionController {
	return &ConnectionController{service: service}
}

// Structure for binding the body of the request
type ConnectionRequest struct {
	ConnectedUserID string `json:"connected_user_id"`
}

func (c *ConnectionController) FollowUser(ctx *gin.Context) {
	userID := ctx.Param("id")

	var request ConnectionRequest
	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Ensure connected_user_id is not empty
	if request.ConnectedUserID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "connected_user_id is required"})
		return
	}

	err := c.service.FollowUser(userID, request.ConnectedUserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Followed successfully"})
}

func (c *ConnectionController) UnfollowUser(ctx *gin.Context) {
	userID := ctx.Param("id")

	var request ConnectionRequest
	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Ensure connected_user_id is not empty
	if request.ConnectedUserID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "connected_user_id is required"})
		return
	}

	err := c.service.UnfollowUser(userID, request.ConnectedUserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Unfollowed successfully"})
}

func (c *ConnectionController) SendFriendRequest(ctx *gin.Context) {
	userID := ctx.Param("id")

	var request ConnectionRequest
	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Ensure connected_user_id is not empty
	if request.ConnectedUserID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "connected_user_id is required"})
		return
	}

	err := c.service.SendFriendRequest(userID, request.ConnectedUserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Friend request sent"})
}
