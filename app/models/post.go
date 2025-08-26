package models

// This struct contain all needed data to render a post.
//
// For posts with attachements this struct should be used in pair with some of attachment structs.
type Post struct {
	ID             int    `json:"id"`
	Content        string `json:"content"`
	AttachmentType string `json:"attachment_type"`
	NumOfComments  int    `json:"num_of_comments"`
	NumOfLikes     int    `json:"num_of_likes"`
	IsLiked        bool   `json:"is_liked"`
	CreatedAt      string `json:"created_at"`
}

// This struct match DB representation of Post. For internal use only
type PostDB struct {
	ID              int    `json:"id"`
	Content         string `json:"content"`
	AttachementType string `json:"attachment_type"`
	CreatedAt       string `json:"created_at"`
}

// This struct purpose is to store single image post's attachment data to Render Image.
type AttachementImage struct {
	ID            int    `json:"id"`
	ImageFilename string `json:"img_filename"`
	CreatedAt     string `json:"created_at"`
}

// This struct purpose is to store multiple images post's attachment data to render Carousel.
type AttachementCarousel struct {
	ID                int    `json:"id"`
	LastImageFilename string `json:"last_img_filename"`
	CreatedAt         string `json:"created_at"`
}

// This struct purpose is to store YouTube Video data to render Video Link.
type AttachementYoutubeVid struct {
	ID          int    `json:"id"`
	VideoId     string `json:"video_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
}
