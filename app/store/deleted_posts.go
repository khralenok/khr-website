package store

import "github.com/khralenok/khr-website/db"

// Delete post with specific ID
func DeletePost(id int) error {
	query := "INSERT INTO deleted_posts(id) VALUES ($1)"

	_, err := db.DB.Exec(query, id)

	if err != nil {
		return err
	}

	return nil
}
