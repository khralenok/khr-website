package models

type Comment struct {
	ID            int    `json:"id"`
	Content       string `json:"content"`
	PostId        int    `json:"post_id"`
	CommentatorId int    `json:"commentator_id"`
	CreatedAt     string `json:"created_at"`
	NumOfReplies  int    `json:"num_of_replies"`
}
