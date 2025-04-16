package models

import "time"

type Connection struct {
	ID              string    `json:"id"`
	UserID          string    `json:"user_id"`
	ConnectedUserID string    `json:"connected_user_id"`
	Status          string    `json:"status"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
