package services

import (
	"github.com/sagar-rathod-devops/do-host-network/internal/models"
	"github.com/sagar-rathod-devops/do-host-network/internal/repositories"
)

type FeedService interface {
	GetFeedChronological(limit int) ([]models.Post, error)
	GetFeedTrending(limit int) ([]models.Post, error)
}

type feedService struct {
	postRepo repositories.PostRepository
}

// NewFeedService creates a new FeedService
func NewFeedService(postRepo repositories.PostRepository) FeedService {
	return &feedService{postRepo: postRepo}
}

// GetFeedChronological gets posts in chronological order
func (s *feedService) GetFeedChronological(limit int) ([]models.Post, error) {
	return s.postRepo.GetPostsChronological(limit)
}

// GetFeedTrending gets trending posts
func (s *feedService) GetFeedTrending(limit int) ([]models.Post, error) {
	return s.postRepo.GetTrendingPosts(limit)
}
