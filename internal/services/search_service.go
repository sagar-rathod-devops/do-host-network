package services

import (
	"github.com/sagar-rathod-devops/do-host-network/internal/models"
	"github.com/sagar-rathod-devops/do-host-network/internal/repositories"
)

type SearchService interface {
	Search(query string, limit int) (models.SearchResult, error)
}

type searchService struct {
	searchRepo repositories.SearchRepository
}

func NewSearchService(searchRepo repositories.SearchRepository) SearchService {
	return &searchService{searchRepo: searchRepo}
}

func (s *searchService) Search(query string, limit int) (models.SearchResult, error) {
	// Use the unified Search method from the repository
	result, err := s.searchRepo.Search(query, limit)
	if err != nil {
		return models.SearchResult{}, err
	}
	return result, nil
}
