package services

import (
	"context"
	"errors"

	"github.com/sagar-rathod-devops/do-host-network/internal/models"
	"github.com/sagar-rathod-devops/do-host-network/internal/repositories"
)

type PostService struct {
	Repo *repositories.PostRepository
}

func NewPostService(repo *repositories.PostRepository) *PostService {
	return &PostService{Repo: repo}
}

func (s *PostService) CreatePost(ctx context.Context, post *models.Post) error {
	if post.Content == "" {
		return errors.New("content cannot be empty")
	}
	return s.Repo.CreatePost(ctx, post)
}

func (s *PostService) GetPostByID(ctx context.Context, id string) (*models.Post, error) {
	return s.Repo.GetPostByID(ctx, id)
}

func (s *PostService) DeletePost(ctx context.Context, id string) error {
	return s.Repo.DeletePost(ctx, id)
}
