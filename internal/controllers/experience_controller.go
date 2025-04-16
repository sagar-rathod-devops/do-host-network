package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sagar-rathod-devops/do-host-network/internal/models"
	"github.com/sagar-rathod-devops/do-host-network/internal/services"
)

type ExperienceController struct {
	service services.ExperienceService
}

func NewExperienceController(service services.ExperienceService) *ExperienceController {
	return &ExperienceController{service: service}
}

func (c *ExperienceController) CreateExperience(ctx *gin.Context) {
	var exp models.Experience
	if err := ctx.ShouldBindJSON(&exp); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.service.CreateExperience(&exp); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, exp)
}

func (c *ExperienceController) GetExperienceByID(ctx *gin.Context) {
	id := ctx.Param("id")
	exp, err := c.service.GetExperienceByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, exp)
}

func (c *ExperienceController) DeleteExperience(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := c.service.DeleteExperience(id); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func (c *ExperienceController) GetAllExperiencesByUserID(ctx *gin.Context) {
	userID := ctx.Param("user_id")
	experiences, err := c.service.GetAllExperiencesByUserID(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, experiences)
}
