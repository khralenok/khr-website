package store

import (
	"database/sql"
	"time"

	"github.com/khralenok/khr-website/db"
	"github.com/khralenok/khr-website/models"
)

// This function return an array of comments for specific post ordered from latest to oldest
func GetComments(postId int) ([]models.Comment, error) {
	var comments []models.Comment

	query := "SELECT * FROM comments WHERE post_id = $1 ORDER BY created_at DESC"

	rows, err := db.DB.Query(query, postId)

	if err == sql.ErrNoRows {
		return []models.Comment{}, nil
	}

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

// This function return a comment with specified ID
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

// This function insert new comment to database
func AddComment(content string, postId int) error {
	query := "INSERT INTO comments(content, post_id) VALUES ($1, $2)"

	_, err := db.DB.Exec(query, content, postId)

	if err != nil {
		return err
	}

	return nil
}

// This function return total amount of comments which belongs to the particular post
func CountPostComments(postID int) int {
	var numOfComments int

	query := "SELECT COUNT(*) FROM comments WHERE post_id = $1"

	err := db.DB.QueryRow(query, postID).Scan(&numOfComments)

	if err != nil {
		return 0
	}

	return numOfComments
}
