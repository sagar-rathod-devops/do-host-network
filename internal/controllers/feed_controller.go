package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sagar-rathod-devops/do-host-network/internal/services"
)

type FeedController struct {
	feedService services.FeedService
}

// NewFeedController creates a new FeedController
func NewFeedController(feedService services.FeedService) *FeedController {
	return &FeedController{feedService: feedService}
}

// GetFeedChronological handles the request for chronological feed
func (c *FeedController) GetFeedChronological(ctx *gin.Context) {
	limit, err := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit"})
		return
	}

	posts, err := c.feedService.GetFeedChronological(limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, posts)
}

// GetFeedTrending handles the request for trending feed
func (c *FeedController) GetFeedTrending(ctx *gin.Context) {
	limit, err := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit"})
		return
	}

	posts, err := c.feedService.GetFeedTrending(limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, posts)
}
