package store

import (
	"database/sql"
	"time"

	"github.com/khralenok/khr-website/db"
	"github.com/khralenok/khr-website/models"
)

// This function return an array of replies for specific comment ordered from latest to oldest
func GetReplies(commentId int) ([]models.Reply, error) {
	var replies []models.Reply

	query := "SELECT r.*, u.username FROM replies r JOIN users u ON r.commentator_id = u.id WHERE r.comment_id = $1 AND NOT EXISTS (SELECT 1 FROM deleted_replies d WHERE d.id = r.id) ORDER BY r.created_at DESC"

	rows, err := db.DB.Query(query, commentId)

	if err == sql.ErrNoRows {
		return []models.Reply{}, nil
	}

	if err != nil {
		return []models.Reply{}, err
	}

	defer rows.Close()

	for rows.Next() {
		nextReply, err := newReply(rows)

		if err != nil {
			return []models.Reply{}, err
		}

		replies = append(replies, nextReply)
	}

	return replies, nil
}

// This function return a reply with specified ID
func GetReply(replyId int) (models.Reply, error) {
	var comment models.Reply

	query := "SELECT r.*, u.username FROM replies r JOIN users u ON r.commentator_id = u.id WHERE r.id=$1 AND NOT EXISTS (SELECT 1 FROM deleted_replies d WHERE d.id = r.id)"

	rows, err := db.DB.Query(query, replyId)

	if err != nil {
		return models.Reply{}, err
	}

	defer rows.Close()

	for rows.Next() {
		comment, err = newReply(rows)

		if err == sql.ErrNoRows {
			break
		}

		if err != nil {
			return models.Reply{}, err
		}
	}

	return comment, nil
}

// This function insert new comment to database
func AddReply(content string, commentId, commentatorId int) error {
	query := "INSERT INTO replies(content, comment_id, commentator_id) VALUES ($1, $2, $3)"

	_, err := db.DB.Exec(query, content, commentId, commentatorId)

	if err != nil {
		return err
	}

	return nil
}

// This function return total amount of comments which belongs to the particular post
func CountPostReplies(postID int) int {
	var numOfReplies int

	query := "SELECT COUNT(*) FROM replies r INNER JOIN comments c ON r.comment_id = c.id WHERE c.post_id = $1 AND NOT EXISTS (SELECT 1 FROM deleted_replies d WHERE d.id = r.id)"

	err := db.DB.QueryRow(query, postID).Scan(&numOfReplies)

	if err != nil {
		return 0
	}

	return numOfReplies
}

// This function return total amount of comments which belongs to the particular post
func CountCommentReplies(commentID int) int {
	var numOfReplies int

	query := "SELECT COUNT(*) FROM replies r INNER JOIN comments c ON r.comment_id = c.id WHERE c.id = $1 AND NOT EXISTS (SELECT 1 FROM deleted_comments d WHERE d.id = c.id)"

	err := db.DB.QueryRow(query, commentID).Scan(&numOfReplies)

	if err != nil {
		return 0
	}

	return numOfReplies
}

// This function make a record about deleted reply into deleted_posts table
func DeleteReply(id int) error {
	query := "INSERT INTO deleted_replies(id) VALUES ($1)"

	_, err := db.DB.Exec(query, id)

	if err != nil {
		return err
	}

	return nil
}

// This function delete all replies belonging to specific comment
func DeleteReplies(commentId int) error {
	query := "SELECT id FROM replies r WHERE r.comment_id = $1 AND NOT EXISTS (SELECT 1 FROM deleted_replies d WHERE d.id = r.id)"

	rows, err := db.DB.Query(query, commentId)

	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {
		var replyId int

		err := rows.Scan(&replyId)

		if err != nil {
			return err
		}

		err = DeleteReply(replyId)

		if err != nil {
			return err
		}
	}

	return nil
}

// This function construct new Post struct from a result of database query
func newReply(row *sql.Rows) (models.Reply, error) {
	var newReply models.Reply
	var rawTime time.Time

	err := row.Scan(&newReply.ID, &newReply.Content, &newReply.CommentId, &newReply.CommentatorId, &rawTime, &newReply.CommentatorUsername)

	if err != nil {
		return models.Reply{}, err
	}

	newReply.CreatedAt = rawTime.Format("02 Jan 2006 15:04")

	return newReply, nil
}
