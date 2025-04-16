package models

import "time"

type Experience struct {
	ID               string     `json:"id"`
	UserID           string     `json:"user_id"`
	CompanyName      string     `json:"company_name"`
	JobTitle         string     `json:"job_title"`
	Location         *string    `json:"location,omitempty"`
	StartDate        *time.Time `json:"start_date,omitempty"` // No time part needed
	EndDate          *time.Time `json:"end_date,omitempty"`   // No time part needed
	CurrentlyWorking bool       `json:"currently_working"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`
}
