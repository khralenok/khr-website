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

func GetVideoAttachment(postId int) (models.AttachementYoutubeVid, error) {
	var attachment models.AttachementYoutubeVid

	query := "SELECT * FROM attachment_youtube_vids WHERE id=$1"

	err := db.DB.QueryRow(query, postId).Scan(&attachment.ID, &attachment.VideoId, &attachment.Title, &attachment.Description, &attachment.CreatedAt)

	if err != nil {
		return models.AttachementYoutubeVid{}, err
	}

	return attachment, nil
}

func AddImageAttachment(postId int, filename string) (models.AttachementImage, error) {
	var newAttachment models.AttachementImage

	query := "INSERT INTO attachment_images(id, img_filename) VALUES ($1, $2) RETURNING *"

	err := db.DB.QueryRow(query, postId, filename).Scan(&newAttachment.ID, &newAttachment.ImageFilename, &newAttachment.CreatedAt)

	if err != nil {
		return models.AttachementImage{}, err
	}

	return newAttachment, nil
}

func GetImageAttachment(postId int) (models.AttachementImage, error) {
	var attachment models.AttachementImage

	query := "SELECT * FROM attachment_images WHERE id=$1"

	err := db.DB.QueryRow(query, postId).Scan(&attachment.ID, &attachment.ImageFilename, &attachment.CreatedAt)

	if err != nil {
		return models.AttachementImage{}, err
	}

	return attachment, nil
}
