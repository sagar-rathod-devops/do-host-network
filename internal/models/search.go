package models

type SearchResult struct {
	Users  []User  `json:"users,omitempty"`  // Reuse existing User model
	Posts  []Post  `json:"posts,omitempty"`  // Reuse existing Post model
	Groups []Group `json:"groups,omitempty"` // Reuse existing Group model
}
