package models

import "time"

type Post struct {
	ID        int64     `json:"id"`
	Content   string    `json:"content"`
	ImageURL  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
}
