package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sagar-rathod-devops/do-host-network/internal/models"
	"github.com/sagar-rathod-devops/do-host-network/internal/services"
)

type CommentController struct {
	Service *services.CommentService
}

type ReactionController struct {
	Service *services.ReactionService
}

func NewCommentController(service *services.CommentService) *CommentController {
	return &CommentController{Service: service}
}

func NewReactionController(service *services.ReactionService) *ReactionController {
	return &ReactionController{Service: service}
}

func (c *CommentController) CreateComment(ctx *gin.Context) {
	postID := ctx.Param("id") // Extract post_id from the URL
	if postID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "post_id is required"})
		return
	}

	var comment models.Comment
	if err := ctx.ShouldBindJSON(&comment); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Assign the post_id from the URL to the comment
	comment.PostID = postID

	// Call the service to create the comment
	if err := c.Service.CreateComment(ctx, &comment); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Comment added successfully"})
}

func (c *ReactionController) AddReaction(ctx *gin.Context) {
	postID := ctx.Param("id")                  // Extract post_id from URL
	reactionType := ctx.Param("reaction_type") // Extract reaction_type from URL

	if postID == "" || reactionType == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "post_id and reaction_type are required"})
		return
	}

	var reaction models.Reaction
	if err := ctx.ShouldBindJSON(&reaction); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if reaction.UserID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "user_id is required"})
		return
	}

	// Set post_id and reaction_type in the reaction object
	reaction.PostID = postID
	reaction.ReactionType = reactionType

	// Call the service to add the reaction
	if err := c.Service.AddReaction(ctx, &reaction); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Reaction added successfully"})
}

func (r *ReactionController) RemoveReaction(ctx *gin.Context) {
	postID := ctx.Param("id")                  // Extract post_id from URL
	reactionType := ctx.Param("reaction_type") // Extract reaction_type from URL
	userID := ctx.Param("user_id")             // Extract user_id from URL

	if postID == "" || reactionType == "" || userID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "post_id, reaction_type, and user_id are required"})
		return
	}

	// Call the service to remove the reaction
	if err := r.Service.RemoveReaction(ctx, postID, userID, reactionType); err != nil {
		if err.Error() == "reaction not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Reaction removed successfully"})
}
