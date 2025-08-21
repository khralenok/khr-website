package models

type Reply struct {
	ID                  int    `json:"id"`
	Content             string `json:"content"`
	CommentatorId       int    `json:"commentator_id"`
	CommentatorUsername string `json:"commentator_username"`
	CommentatorAvatar   string `json:"commentator_avatar"`
	CommentId           int    `json:"comment_id"`
	CreatedAt           string `json:"created_at"`
}
