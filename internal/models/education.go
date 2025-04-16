package models

type Education struct {
	ID           string  `json:"id"`
	UserID       string  `json:"user_id"`
	SchoolName   string  `json:"school_name"`
	Degree       *string `json:"degree,omitempty"`
	FieldOfStudy *string `json:"field_of_study,omitempty"`
	StartDate    *string `json:"start_date,omitempty"`
	EndDate      *string `json:"end_date,omitempty"`
	CreatedAt    string  `json:"created_at"`
	UpdatedAt    string  `json:"updated_at"`
}
