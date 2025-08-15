package store

import "github.com/khralenok/khr-website/db"

// This function insert deleted post to deleted_posts table
func DeleteComment(id int) error {
	query := "INSERT INTO deleted_comments(id) VALUES ($1)"

	_, err := db.DB.Exec(query, id)

	if err != nil {
		return err
	}

	return nil
}
