package repositories

import (
	"database/sql"

	"github.com/sagar-rathod-devops/do-host-network/internal/models"
)

type SearchRepository interface {
	Search(query string, limit int) (models.SearchResult, error)
}

type searchRepository struct {
	db *sql.DB
}

func NewSearchRepository(db *sql.DB) SearchRepository {
	return &searchRepository{db: db}
}

func (r *searchRepository) Search(query string, limit int) (models.SearchResult, error) {
	var result models.SearchResult

	// Search Users
	userQuery := `SELECT id, email, username FROM users WHERE username ILIKE $1 OR email ILIKE $1 LIMIT $2`
	userRows, err := r.db.Query(userQuery, "%"+query+"%", limit)
	if err != nil {
		return result, err
	}
	defer userRows.Close()

	for userRows.Next() {
		var user models.User
		if err := userRows.Scan(&user.ID, &user.Email, &user.Username); err != nil {
			return result, err
		}
		result.Users = append(result.Users, user)
	}

	// Search Posts
	postQuery := `SELECT id, user_id, content FROM posts WHERE content ILIKE $1 LIMIT $2`
	postRows, err := r.db.Query(postQuery, "%"+query+"%", limit)
	if err != nil {
		return result, err
	}
	defer postRows.Close()

	for postRows.Next() {
		var post models.Post
		if err := postRows.Scan(&post.ID, &post.UserID, &post.Content); err != nil {
			return result, err
		}
		result.Posts = append(result.Posts, post)
	}

	// Search Groups
	groupQuery := `SELECT id, name, description FROM groups WHERE name ILIKE $1 OR description ILIKE $1 LIMIT $2`
	groupRows, err := r.db.Query(groupQuery, "%"+query+"%", limit)
	if err != nil {
		return result, err
	}
	defer groupRows.Close()

	for groupRows.Next() {
		var group models.Group
		if err := groupRows.Scan(&group.ID, &group.Name, &group.Description); err != nil {
			return result, err
		}
		result.Groups = append(result.Groups, group)
	}

	return result, nil
}
