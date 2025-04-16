// user_profiles_controller.go
package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sagar-rathod-devops/do-host-network/internal/models"
	"github.com/sagar-rathod-devops/do-host-network/internal/services"
)

type UserProfileController struct {
	Service services.UserProfileService
}

func NewUserProfileController(service services.UserProfileService) *UserProfileController {
	return &UserProfileController{Service: service}
}

func (c *UserProfileController) GetUserProfile(ctx *gin.Context) {
	id := ctx.Param("id")
	profile, email, err := c.Service.GetUserProfile(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// Create a response with the email and profile details
	response := gin.H{
		"id":                  profile.ID,
		"user_id":             profile.UserID,
		"email":               email,
		"first_name":          profile.FirstName,
		"last_name":           profile.LastName,
		"bio":                 profile.Bio,
		"profile_picture_url": profile.ProfilePictureURL,
		"location":            profile.Location,
		"birth_date":          profile.BirthDate,
		"headline":            profile.Headline,
		"industry":            profile.Industry,
		"created_at":          profile.CreatedAt,
		"updated_at":          profile.UpdatedAt,
	}

	ctx.JSON(http.StatusOK, response)
}

func (c *UserProfileController) CreateOrUpdateUserProfile(ctx *gin.Context) {
	var profile models.UserProfile
	if err := ctx.ShouldBindJSON(&profile); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	updatedProfile, err := c.Service.CreateOrUpdateUserProfile(&profile)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, updatedProfile)
}

func (c *UserProfileController) DeleteUserProfile(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := c.Service.DeleteUserProfile(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}
