package models

import (
	"time"
)

type MediaFile struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	MediaURL  string    `json:"media_url"`
	MediaType string    `json:"media_type"`
	CreatedAt time.Time `json:"created_at"`
}
