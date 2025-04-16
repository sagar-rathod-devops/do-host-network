package repositories

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/sagar-rathod-devops/do-host-network/internal/models"
)

type PostRepository struct {
	DB *sql.DB
}

func NewPostRepository(db *sql.DB) *PostRepository {
	return &PostRepository{DB: db}
}

// CreatePost inserts a new post into the database
func (r *PostRepository) CreatePost(ctx context.Context, post *models.Post) error {
	// Check if the user_id exists in the users table
	var exists bool
	checkUserQuery := `SELECT EXISTS(SELECT 1 FROM users WHERE id = $1)`
	err := r.DB.QueryRowContext(ctx, checkUserQuery, post.UserID).Scan(&exists)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("user_id does not exist")
	}

	// Insert the post
	query := `INSERT INTO posts (id, user_id, content, media_url)
			  VALUES ($1, $2, $3, $4)`
	_, err = r.DB.ExecContext(ctx, query, post.ID, post.UserID, post.Content, post.MediaURL)
	return err
}

// GetPostByID retrieves a post by its ID
func (r *PostRepository) GetPostByID(ctx context.Context, id string) (*models.Post, error) {
	query := `SELECT id, user_id, content, media_url, created_at, updated_at FROM posts WHERE id = $1`
	row := r.DB.QueryRowContext(ctx, query, id)

	var post models.Post
	if err := row.Scan(&post.ID, &post.UserID, &post.Content, &post.MediaURL, &post.CreatedAt, &post.UpdatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("post not found")
		}
		return nil, err
	}

	return &post, nil
}

// DeletePost deletes a post by its ID
func (r *PostRepository) DeletePost(ctx context.Context, id string) error {
	query := `DELETE FROM posts WHERE id = $1`
	result, err := r.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	// Ensure a post was deleted
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("post not found")
	}

	return nil
}

// GetPostsChronological fetches posts in chronological order
func (r *PostRepository) GetPostsChronological(limit int) ([]models.Post, error) {
	query := `SELECT id, user_id, content, media_url, created_at, updated_at 
              FROM posts 
              ORDER BY created_at DESC 
              LIMIT $1`
	rows, err := r.DB.Query(query, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		if err := rows.Scan(&post.ID, &post.UserID, &post.Content, &post.MediaURL, &post.CreatedAt, &post.UpdatedAt); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

// GetTrendingPosts fetches trending posts based on your criteria (placeholder logic)
func (r *PostRepository) GetTrendingPosts(limit int) ([]models.Post, error) {
	query := `SELECT id, user_id, content, media_url, created_at, updated_at 
              FROM posts 
              WHERE created_at > $1
              ORDER BY random() 
              LIMIT $2` // Example logic for trending
	oneWeekAgo := time.Now().AddDate(0, 0, -7)
	rows, err := r.DB.Query(query, oneWeekAgo, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		if err := rows.Scan(&post.ID, &post.UserID, &post.Content, &post.MediaURL, &post.CreatedAt, &post.UpdatedAt); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}
