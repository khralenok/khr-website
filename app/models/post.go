package models

type Post struct {
	ID            int    `json:"id"`
	Content       string `json:"content"`
	AttachmentKey string `json:"attachment_key"` // if empty => no attachements, if looks like "oCj18Uv-zcY" means attachment is YouTube link, and if looks like "post_id-illustration_id" means attachment is images
	CreatedAt     string `json:"created_at"`
	NumOfComments int    `json:"num_of_comments"`
	NumOfLikes    int    `json:"num_of_likes"`
	IsLiked       bool   `json:"is_liked"`
}

type VideoAttachement struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}
