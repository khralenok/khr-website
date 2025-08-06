package store

import "github.com/khralenok/khr-website/db"

// This function insert deleted post to deleted_posts table
func DeletePost(id int) error {
	query := "INSERT INTO deleted_posts(id) VALUES ($1)"

	_, err := db.DB.Exec(query, id)

	if err != nil {
		return err
	}

	return nil
}
