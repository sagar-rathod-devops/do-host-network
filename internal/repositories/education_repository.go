package repositories

import (
	"database/sql"
	"errors"

	"github.com/sagar-rathod-devops/do-host-network/internal/models"
)

type EducationRepository struct {
	DB *sql.DB
}

// GetEducationByID retrieves an education record by its ID
func (r *EducationRepository) GetEducationByID(id string) (*models.Education, error) {
	query := `
		SELECT id, user_id, school_name, degree, field_of_study, start_date, end_date, created_at, updated_at
		FROM education
		WHERE id = $1
	`

	edu := &models.Education{}
	row := r.DB.QueryRow(query, id)
	err := row.Scan(&edu.ID, &edu.UserID, &edu.SchoolName, &edu.Degree, &edu.FieldOfStudy, &edu.StartDate, &edu.EndDate, &edu.CreatedAt, &edu.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // No record found
		}
		return nil, err // Other errors
	}

	return edu, nil
}

// CreateEducation inserts a new education record into the database
func (r *EducationRepository) CreateEducation(edu *models.Education) error {
	query := `
		INSERT INTO education (user_id, school_name, degree, field_of_study, start_date, end_date)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, created_at, updated_at
	`

	err := r.DB.QueryRow(
		query,
		edu.UserID,
		edu.SchoolName,
		edu.Degree,
		edu.FieldOfStudy,
		edu.StartDate,
		edu.EndDate,
	).Scan(&edu.ID, &edu.CreatedAt, &edu.UpdatedAt)

	if err != nil {
		return err
	}

	return nil
}

// DeleteEducationByID deletes an education record by its ID
func (r *EducationRepository) DeleteEducationByID(id string) error {
	query := `
		DELETE FROM education
		WHERE id = $1
	`

	result, err := r.DB.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return sql.ErrNoRows // No record found to delete
	}

	return nil
}

// ListEducationByUserID retrieves all education records for a given user
func (r *EducationRepository) ListEducationByUserID(userID string) ([]*models.Education, error) {
	query := `
		SELECT id, user_id, school_name, degree, field_of_study, start_date, end_date, created_at, updated_at
		FROM education
		WHERE user_id = $1
	`

	rows, err := r.DB.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var educationList []*models.Education
	for rows.Next() {
		edu := &models.Education{}
		err := rows.Scan(&edu.ID, &edu.UserID, &edu.SchoolName, &edu.Degree, &edu.FieldOfStudy, &edu.StartDate, &edu.EndDate, &edu.CreatedAt, &edu.UpdatedAt)
		if err != nil {
			return nil, err
		}
		educationList = append(educationList, edu)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return educationList, nil
}
