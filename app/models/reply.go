package models

type Reply struct {
	ID            int    `json:"id"`
	Content       string `json:"content"`
	CommentatorId int    `json:"commentator_id"`
	CommentId     int    `json:"comment_id"`
	CreatedAt     string `json:"created_at"`
}
