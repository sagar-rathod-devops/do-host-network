package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sagar-rathod-devops/do-host-network/internal/services"
)

type SearchController struct {
	searchService services.SearchService
}

func NewSearchController(searchService services.SearchService) *SearchController {
	return &SearchController{searchService: searchService}
}

func (c *SearchController) Search(ctx *gin.Context) {
	// Parse the search query parameter
	query := ctx.Query("q")
	if query == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Search query cannot be empty"})
		return
	}

	// Parse the limit parameter with a default value
	limit, err := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit"})
		return
	}

	// Fetch search results using the service
	result, err := c.searchService.Search(query, limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the search results
	ctx.JSON(http.StatusOK, result)
}
