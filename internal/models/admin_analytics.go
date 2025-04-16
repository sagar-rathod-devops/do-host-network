package models

import "time"

type AdminLog struct {
	ID         string    `json:"id"`
	AdminID    string    `json:"admin_id"`
	TargetID   string    `json:"target_id"`
	TargetType string    `json:"target_type"`
	Action     string    `json:"action"`
	CreatedAt  time.Time `json:"created_at"`
}

type PostInteraction struct {
	ID        string    `json:"id"`
	PostID    string    `json:"post_id"`
	Likes     int       `json:"likes"`
	Comments  int       `json:"comments"`
	CreatedAt time.Time `json:"created_at"`
}

type UserAnalytics struct {
	ID              string    `json:"id"`
	UserID          string    `json:"user_id"`
	EngagementScore int       `json:"engagement_score"`
	Active          bool      `json:"active"`
	CreatedAt       time.Time `json:"created_at"`
}
