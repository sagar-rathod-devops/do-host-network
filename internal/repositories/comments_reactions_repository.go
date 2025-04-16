package repositories

import (
	"context"
	"database/sql"
	"errors"

	"github.com/sagar-rathod-devops/do-host-network/internal/models"
)

type CommentRepository struct {
	DB *sql.DB
}

type ReactionRepository struct {
	DB *sql.DB
}

func NewCommentRepository(db *sql.DB) *CommentRepository {
	return &CommentRepository{DB: db}
}

func NewReactionRepository(db *sql.DB) *ReactionRepository {
	return &ReactionRepository{DB: db}
}

func (r *CommentRepository) CreateComment(ctx context.Context, comment *models.Comment) error {
	query := `INSERT INTO comments (post_id, user_id, content, created_at, updated_at)
              VALUES ($1, $2, $3, NOW(), NOW())`
	_, err := r.DB.ExecContext(ctx, query, comment.PostID, comment.UserID, comment.Content)
	return err
}

func (r *ReactionRepository) AddReaction(ctx context.Context, reaction *models.Reaction) error {
	// Check if the user has already reacted to this post
	var count int
	err := r.DB.QueryRowContext(ctx, `SELECT COUNT(*) FROM reactions WHERE post_id = $1 AND user_id = $2`, reaction.PostID, reaction.UserID).Scan(&count)
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("user has already reacted to this post")
	}

	// Insert the new reaction
	query := `INSERT INTO reactions (post_id, user_id, reaction_type, created_at)
              VALUES ($1, $2, $3, NOW())`
	_, err = r.DB.ExecContext(ctx, query, reaction.PostID, reaction.UserID, reaction.ReactionType)
	return err
}

func (r *ReactionRepository) RemoveReaction(ctx context.Context, postID, userID, reactionType string) error {
	// Check if the reaction exists
	var count int
	err := r.DB.QueryRowContext(ctx, `SELECT COUNT(*) FROM reactions WHERE post_id = $1 AND user_id = $2 AND reaction_type = $3`, postID, userID, reactionType).Scan(&count)
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("reaction not found")
	}

	// Delete the reaction
	query := `DELETE FROM reactions WHERE post_id = $1 AND user_id = $2 AND reaction_type = $3`
	_, err = r.DB.ExecContext(ctx, query, postID, userID, reactionType)
	return err
}
