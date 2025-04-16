package repositories

import (
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/sagar-rathod-devops/do-host-network/internal/models"
)

type UserProfileRepository struct {
	DB *sql.DB // PostgreSQL database connection
}

func NewUserProfileRepository(db *sql.DB) *UserProfileRepository {
	return &UserProfileRepository{DB: db}
}

// GetByID fetches a user profile by ID and includes the email from the users table
func (r *UserProfileRepository) GetByID(id string) (*models.UserProfile, string, error) {
	query := `SELECT up.id, up.user_id, u.email, up.first_name, up.last_name, up.bio, 
                     up.profile_picture_url, up.location, up.birth_date, 
                     up.headline, up.industry, up.created_at, up.updated_at
              FROM user_profiles up
              JOIN users u ON up.user_id = u.id
              WHERE up.id = $1`

	row := r.DB.QueryRow(query, id)
	profile := &models.UserProfile{}
	var email string

	err := row.Scan(
		&profile.ID, &profile.UserID, &email, &profile.FirstName, &profile.LastName, &profile.Bio,
		&profile.ProfilePictureURL, &profile.Location, &profile.BirthDate, &profile.Headline,
		&profile.Industry, &profile.CreatedAt, &profile.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, "", errors.New("user profile not found")
		}
		return nil, "", err
	}

	return profile, email, nil
}

// Create inserts a new user profile into the database
func (r *UserProfileRepository) Create(profile *models.UserProfile) (*models.UserProfile, error) {
	// Check if the user_id exists in the users table
	existsQuery := `SELECT COUNT(1) FROM users WHERE id = $1`
	var count int
	err := r.DB.QueryRow(existsQuery, profile.UserID).Scan(&count)
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return nil, errors.New("user_id does not exist in users table")
	}

	query := `INSERT INTO user_profiles (id, user_id, first_name, last_name, bio, 
                                         profile_picture_url, location, birth_date, 
                                         headline, industry, created_at, updated_at)
              VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, NOW(), NOW())
              RETURNING id, created_at, updated_at`

	profile.ID = uuid.New().String() // Generate a new UUID if not provided
	err = r.DB.QueryRow(
		query,
		profile.ID, profile.UserID, profile.FirstName, profile.LastName, profile.Bio,
		profile.ProfilePictureURL, profile.Location, profile.BirthDate, profile.Headline, profile.Industry,
	).Scan(&profile.ID, &profile.CreatedAt, &profile.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return profile, nil
}

// Delete removes a user profile from the database by ID
func (r *UserProfileRepository) Delete(id string) error {
	query := `DELETE FROM user_profiles WHERE id = $1`
	result, err := r.DB.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("user profile not found")
	}

	return nil
}
