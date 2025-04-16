package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sagar-rathod-devops/do-host-network/internal/models"
	"github.com/sagar-rathod-devops/do-host-network/internal/services"
)

type MediaController struct {
	Service *services.MediaService
}

func NewMediaController(service *services.MediaService) *MediaController {
	return &MediaController{Service: service}
}

func (c *MediaController) CreateMedia(ctx *gin.Context) {
	var media models.MediaFile
	if err := ctx.ShouldBindJSON(&media); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.Service.CreateMedia(&media); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save media"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Media created successfully"})
}

func (c *MediaController) GetMediaByUserID(ctx *gin.Context) {
	userID := ctx.Param("user_id")
	mediaFiles, err := c.Service.GetMediaByUserID(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch media"})
		return
	}

	ctx.JSON(http.StatusOK, mediaFiles)
}
