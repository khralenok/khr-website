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

	query := "SELECT c.*, u.username FROM comments c JOIN users u ON c.commentator_id = u.id WHERE NOT EXISTS (SELECT 1 FROM deleted_comments d WHERE d.id = c.id) AND c.post_id = $1 ORDER BY c.created_at DESC"

	rows, err := db.DB.Query(query, postId)

	if err == sql.ErrNoRows {
		return []models.Comment{}, nil
	}

	if err != nil {
		return []models.Comment{}, err
	}

	defer rows.Close()

	for rows.Next() {
		nextComment, err := newComment(rows)

		if err != nil {
			return []models.Comment{}, err
		}

		comments = append(comments, nextComment)
	}

	return comments, nil
}

// This function return a comment with specified ID
func GetComment(commentId int) (models.Comment, error) {
	var comment models.Comment

	query := "SELECT c.*, u.username FROM comments c JOIN users u ON c.commentator_id = u.id WHERE NOT EXISTS (SELECT 1 FROM deleted_comments d WHERE d.id = c.id) AND c.id=$1"

	rows, err := db.DB.Query(query, commentId)

	if err != nil {
		return models.Comment{}, err
	}

	defer rows.Close()

	for rows.Next() {
		comment, err = newComment(rows)

		if err == sql.ErrNoRows {
			break
		}

		if err != nil {
			return models.Comment{}, err
		}
	}

	return comment, nil
}

// This function insert new comment to database
func AddComment(content string, postId, commentatorId int) error {
	query := "INSERT INTO comments(content, post_id, commentator_id) VALUES ($1, $2, $3)"

	_, err := db.DB.Exec(query, content, postId, commentatorId)

	if err != nil {
		return err
	}

	return nil
}

// This function update comment with specific ID
func UpdateComment(content string, commentId int) error {
	query := "UPDATE comments SET content=$1 WHERE id = $2"

	_, err := db.DB.Exec(query, content, commentId)

	if err != nil {
		return err
	}

	return nil
}

// This function return total amount of comments which belongs to the particular post
func CountPostComments(postID int) int {
	var numOfComments int

	query := "SELECT COUNT(*) FROM comments c WHERE post_id = $1 AND NOT EXISTS (SELECT 1 FROM deleted_comments d WHERE d.id = c.id)"

	err := db.DB.QueryRow(query, postID).Scan(&numOfComments)

	if err != nil {
		return 0
	}

	return numOfComments
}

// This function insert deleted post to deleted_posts table
func DeleteComment(id int) error {
	query := "INSERT INTO deleted_comments(id) VALUES ($1)"

	_, err := db.DB.Exec(query, id)

	if err != nil {
		return err
	}

	return nil
}

// This function construct new Post struct from a result of database query
func newComment(row *sql.Rows) (models.Comment, error) {
	var newComment models.Comment
	var rawTime time.Time

	err := row.Scan(&newComment.ID, &newComment.Content, &newComment.PostId, &newComment.CommentatorId, &rawTime, &newComment.CommentatorUsername)

	if err != nil {
		return models.Comment{}, err
	}

	newComment.NumOfReplies = CountCommentReplies(newComment.ID)

	newComment.CreatedAt = rawTime.Format("02 Jan 2006 15:04")

	return newComment, nil
}
