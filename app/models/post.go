package models

type Post struct {
	ID            int    `json:"id"`
	Content       string `json:"content"`
	ImageURL      string `json:"image_url"`
	CreatedAt     string `json:"created_at"`
	NumOfComments int    `json:"num_of_comments"`
	NumOfLikes    int    `json:"num_of_likes"`
	IsLiked       bool   `json:"is_liked"`
}
