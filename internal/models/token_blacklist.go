package models

import "time"

type TokenBlacklist struct {
	TokenID   string    `json:"token_id"`
	UserID    string    `json:"user_id"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"created_at"`
}
