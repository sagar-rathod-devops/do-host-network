package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sagar-rathod-devops/do-host-network/internal/services"
)

type AdminController struct {
	service services.AdminService
}

type AnalyticsController struct {
	service services.AnalyticsService
}

func NewAdminController(service services.AdminService) *AdminController {
	return &AdminController{service: service}
}

func NewAnalyticsController(service services.AnalyticsService) *AnalyticsController {
	return &AnalyticsController{service: service}
}

// ModeratePost: Moderates the post by its ID
func (c *AdminController) ModeratePost(ctx *gin.Context) {
	// Extract the post ID from the URL
	postID := ctx.Param("id")

	// Validate UUID format for postID
	_, err := uuid.Parse(postID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID format for post ID"})
		return
	}

	// Extract the admin ID from the context (usually from JWT or session)
	adminID := ctx.GetHeader("Admin-ID") // Assuming admin ID is passed in the header (you could also extract from JWT)
	if adminID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Admin ID is missing"})
		return
	}

	// Validate UUID format for adminID
	_, err = uuid.Parse(adminID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID format for admin ID"})
		return
	}

	// Call service to moderate the post
	err = c.service.ModeratePost(postID, adminID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Post moderated successfully"})
}

// BanUser: Bans the user by their ID
func (c *AdminController) BanUser(ctx *gin.Context) {
	// Extract the user ID from the URL
	userID := ctx.Param("id")

	// Validate UUID format for userID
	_, err := uuid.Parse(userID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID format for user ID"})
		return
	}

	// Extract the admin ID from the context (usually from JWT or session)
	adminID := ctx.GetHeader("Admin-ID") // Assuming admin ID is passed in the header (you could also extract from JWT)
	if adminID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Admin ID is missing"})
		return
	}

	// Validate UUID format for adminID
	_, err = uuid.Parse(adminID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID format for admin ID"})
		return
	}

	// Call service to ban the user
	err = c.service.BanUser(userID, adminID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User banned successfully"})
}

func (c *AnalyticsController) GetPostInteractions(ctx *gin.Context) {
	interactions, err := c.service.GetPostInteractions()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, interactions)
}

func (c *AnalyticsController) GetUserAnalytics(ctx *gin.Context) {
	analytics, err := c.service.GetUserAnalytics()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, analytics)
}
