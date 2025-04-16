package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sagar-rathod-devops/do-host-network/internal/models"
	"github.com/sagar-rathod-devops/do-host-network/internal/services"
)

type EducationController struct {
	service *services.EducationService
}

func NewEducationController(service *services.EducationService) *EducationController {
	return &EducationController{service: service}
}

// GetEducation handles GET /education/:id
func (ctrl *EducationController) GetEducation(c *gin.Context) {
	id := c.Param("id")
	edu, err := ctrl.service.GetEducationByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if edu == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Education record not found"})
		return
	}
	c.JSON(http.StatusOK, edu)
}

// CreateEducation handles POST /education
func (ctrl *EducationController) CreateEducation(c *gin.Context) {
	edu := &models.Education{}
	if err := c.ShouldBindJSON(edu); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.service.CreateEducation(edu); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, edu)
}

// DeleteEducation handles DELETE /education/:id
func (ctrl *EducationController) DeleteEducation(c *gin.Context) {
	id := c.Param("id")
	if err := ctrl.service.DeleteEducationByID(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Education record deleted"})
}

// ListEducationByUser handles GET /education/user/:user_id
func (ctrl *EducationController) ListEducationByUser(c *gin.Context) {
	userID := c.Param("user_id")
	educationList, err := ctrl.service.ListEducationByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, educationList)
}
