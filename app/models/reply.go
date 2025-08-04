package models

type Reply struct {
	ID        int64  `json:"id"`
	Content   string `json:"content"`
	CommentId int64  `json:"post_id"`
	CreatedAt string `json:"created_at"`
}
