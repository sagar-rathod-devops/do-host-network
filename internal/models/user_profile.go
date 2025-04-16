// models/user_profile.go
package models

import "time"

type UserProfile struct {
	ID                string    `json:"id"`
	UserID            string    `json:"user_id"`
	FirstName         string    `json:"first_name"`
	LastName          string    `json:"last_name"`
	Bio               string    `json:"bio"`
	ProfilePictureURL string    `json:"profile_picture_url"`
	Location          string    `json:"location"`
	BirthDate         string    `json:"birth_date"`
	Headline          string    `json:"headline"`
	Industry          string    `json:"industry"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}
