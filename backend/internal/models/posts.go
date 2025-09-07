package models

import "time"

type Post struct {
	ID        int       `json:"id"`
	UserID    int       `json:"userId"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt"`
}
