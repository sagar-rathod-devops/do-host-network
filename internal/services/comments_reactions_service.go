package services

import (
	"context"

	"github.com/sagar-rathod-devops/do-host-network/internal/models"
	"github.com/sagar-rathod-devops/do-host-network/internal/repositories"
)

type CommentService struct {
	Repo *repositories.CommentRepository
}

type ReactionService struct {
	Repo *repositories.ReactionRepository
}

func NewCommentService(repo *repositories.CommentRepository) *CommentService {
	return &CommentService{Repo: repo}
}

func NewReactionService(repo *repositories.ReactionRepository) *ReactionService {
	return &ReactionService{Repo: repo}
}

func (s *CommentService) CreateComment(ctx context.Context, comment *models.Comment) error {
	return s.Repo.CreateComment(ctx, comment)
}

func (s *ReactionService) AddReaction(ctx context.Context, reaction *models.Reaction) error {
	return s.Repo.AddReaction(ctx, reaction)
}

func (s *ReactionService) RemoveReaction(ctx context.Context, postID, userID, reactionType string) error {
	return s.Repo.RemoveReaction(ctx, postID, userID, reactionType)
}
