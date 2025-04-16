package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sagar-rathod-devops/do-host-network/internal/models"
	"github.com/sagar-rathod-devops/do-host-network/internal/services"
)

type GroupController struct {
	Service *services.GroupService
}

func NewGroupController(service *services.GroupService) *GroupController {
	return &GroupController{Service: service}
}

func (ctrl *GroupController) CreateGroup(c *gin.Context) {
	var group models.Group
	if err := c.ShouldBindJSON(&group); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := ctrl.Service.CreateGroup(&group)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create group"})
		return
	}
	c.JSON(http.StatusCreated, group)
}

func (ctrl *GroupController) GetGroup(c *gin.Context) {
	id := c.Param("id")

	group, err := ctrl.Service.GetGroupByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Group not found"})
		return
	}
	c.JSON(http.StatusOK, group)
}

func (ctrl *GroupController) JoinGroup(c *gin.Context) {
	var member models.GroupMember
	if err := c.ShouldBindJSON(&member); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	member.GroupID = c.Param("id")
	err := ctrl.Service.AddGroupMember(&member)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to join group"})
		return
	}
	c.JSON(http.StatusCreated, member)
}
