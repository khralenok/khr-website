package store

import (
	"github.com/khralenok/khr-website/db"
	"github.com/khralenok/khr-website/models"
)

func AddVideoAttachment(postId int, videoId, title, description string) (models.AttachementYoutubeVid, error) {
	var newAttachment models.AttachementYoutubeVid

	query := "INSERT INTO attachment_youtube_vids(id, video_id, title, description) VALUES ($1, $2, $3, $4) RETURNING *"

	err := db.DB.QueryRow(query, postId, videoId, title, description).Scan(&newAttachment.ID, &newAttachment.VideoId, &newAttachment.Title, &newAttachment.Description, &newAttachment.CreatedAt)

	if err != nil {
		return models.AttachementYoutubeVid{}, err
	}

	return newAttachment, nil
}
