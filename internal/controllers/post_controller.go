package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sagar-rathod-devops/do-host-network/internal/models"
	"github.com/sagar-rathod-devops/do-host-network/internal/services"
)

type PostController struct {
	Service *services.PostService
}

func NewPostController(service *services.PostService) *PostController {
	return &PostController{Service: service}
}

func (c *PostController) CreatePost(ctx *gin.Context) {
	var post models.Post
	if err := ctx.ShouldBindJSON(&post); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	post.ID = uuid.New()
	if err := c.Service.CreatePost(ctx, &post); err != nil {
		// ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create post"})
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, post)
}

func (c *PostController) GetPost(ctx *gin.Context) {
	id := ctx.Param("id")

	post, err := c.Service.GetPostByID(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	ctx.JSON(http.StatusOK, post)
}

func (c *PostController) DeletePost(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := c.Service.DeletePost(ctx, id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete post"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
}
