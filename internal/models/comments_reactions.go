package models

type Comment struct {
	ID        string `json:"id"`
	PostID    string `json:"post_id"`
	UserID    string `json:"user_id"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type Reaction struct {
	ID           string `json:"id"`
	PostID       string `json:"post_id"`
	UserID       string `json:"user_id"`
	ReactionType string `json:"reaction_type"`
	CreatedAt    string `json:"created_at"`
}
