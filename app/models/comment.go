package models

type Comment struct {
	ID        int64  `json:"id"`
	Content   string `json:"content"`
	PostId    int64  `json:"post_id"`
	CreatedAt string `json:"created_at"`
}
