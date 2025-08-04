package store

import (
	"time"

	"github.com/khralenok/khr-website/db"
	"github.com/khralenok/khr-website/models"
)

// Return array of comments for specific post ordered from latest to oldest
func GetComments(postId int) ([]models.Comment, error) {
	var comments []models.Comment

	query := "SELECT * FROM comments WHERE post_id = $1 ORDER BY created_at DESC"

	rows, err := db.DB.Query(query, postId)

	if err != nil {
		return []models.Comment{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var nextComment models.Comment
		var rawTime time.Time

		err := rows.Scan(&nextComment.ID, &nextComment.Content, &nextComment.PostId, &rawTime)

		if err != nil {
			return []models.Comment{}, err
		}

		nextComment.CreatedAt = rawTime.Format("02 Jan 2006 15:04")

		comments = append(comments, nextComment)
	}

	return comments, nil
}

// Return a comment with specified ID
func GetComment(postId int) (models.Comment, error) {
	var post models.Comment
	var rawTime time.Time

	query := "SELECT id, content, image_url, created_at FROM posts WHERE id=$1"

	err := db.DB.QueryRow(query, postId).Scan(&post.ID, &post.Content, &rawTime)

	if err != nil {
		return models.Comment{}, err
	}

	post.CreatedAt = rawTime.Format("02 Jan 2006 15:04")

	return post, nil
}

// Insert new comment to DB
func AddComment(content string, postId int) error {
	query := "INSERT INTO comments(content, post_id) VALUES ($1, $2)"

	_, err := db.DB.Exec(query, content, postId)

	if err != nil {
		return err
	}

	return nil
}
