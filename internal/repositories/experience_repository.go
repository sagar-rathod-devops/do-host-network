package repositories

import (
	"database/sql"
	"errors"

	"github.com/sagar-rathod-devops/do-host-network/internal/models"
)

type ExperienceRepository interface {
	CreateExperience(exp *models.Experience) error
	GetExperienceByID(id string) (*models.Experience, error)
	DeleteExperience(id string) error
	GetAllExperiencesByUserID(userID string) ([]models.Experience, error)
}

type experienceRepository struct {
	db *sql.DB
}

func NewExperienceRepository(db *sql.DB) ExperienceRepository {
	return &experienceRepository{db: db}
}

func (r *experienceRepository) CreateExperience(exp *models.Experience) error {
	query := `
        INSERT INTO experience (user_id, company_name, job_title, location, start_date, end_date, currently_working)
        VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id, created_at, updated_at
    `
	return r.db.QueryRow(query, exp.UserID, exp.CompanyName, exp.JobTitle, exp.Location, exp.StartDate, exp.EndDate, exp.CurrentlyWorking).
		Scan(&exp.ID, &exp.CreatedAt, &exp.UpdatedAt)
}

func (r *experienceRepository) GetExperienceByID(id string) (*models.Experience, error) {
	query := `SELECT * FROM experience WHERE id = $1`
	exp := &models.Experience{}
	err := r.db.QueryRow(query, id).Scan(
		&exp.ID, &exp.UserID, &exp.CompanyName, &exp.JobTitle, &exp.Location,
		&exp.StartDate, &exp.EndDate, &exp.CurrentlyWorking, &exp.CreatedAt, &exp.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, errors.New("experience not found")
	}
	return exp, err
}

func (r *experienceRepository) DeleteExperience(id string) error {
	query := `DELETE FROM experience WHERE id = $1`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		return errors.New("experience not found")
	}
	return nil
}

func (r *experienceRepository) GetAllExperiencesByUserID(userID string) ([]models.Experience, error) {
	query := `SELECT * FROM experience WHERE user_id = $1`
	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	experiences := []models.Experience{}
	for rows.Next() {
		exp := models.Experience{}
		err := rows.Scan(
			&exp.ID, &exp.UserID, &exp.CompanyName, &exp.JobTitle, &exp.Location,
			&exp.StartDate, &exp.EndDate, &exp.CurrentlyWorking, &exp.CreatedAt, &exp.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		experiences = append(experiences, exp)
	}
	return experiences, nil
}
