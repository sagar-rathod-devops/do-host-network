package repositories

import (
	"database/sql"

	"github.com/sagar-rathod-devops/do-host-network/internal/models"
)

type MediaRepository struct {
	DB *sql.DB
}

func NewMediaRepository(db *sql.DB) *MediaRepository {
	return &MediaRepository{DB: db}
}

func (r *MediaRepository) CreateMedia(media *models.MediaFile) error {
	query := `INSERT INTO media_files (id, user_id, media_url, media_type, created_at) 
              VALUES ($1, $2, $3, $4, $5)`
	_, err := r.DB.Exec(query, media.ID, media.UserID, media.MediaURL, media.MediaType, media.CreatedAt)
	return err
}

func (r *MediaRepository) GetMediaByUserID(userID string) ([]models.MediaFile, error) {
	query := `SELECT id, user_id, media_url, media_type, created_at 
              FROM media_files WHERE user_id = $1`
	rows, err := r.DB.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var mediaFiles []models.MediaFile
	for rows.Next() {
		var media models.MediaFile
		if err := rows.Scan(&media.ID, &media.UserID, &media.MediaURL, &media.MediaType, &media.CreatedAt); err != nil {
			return nil, err
		}
		mediaFiles = append(mediaFiles, media)
	}
	return mediaFiles, nil
}
