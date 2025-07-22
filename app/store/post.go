package store

import (
	"github.com/khralenok/khr-website/db"
	"github.com/khralenok/khr-website/models"
)

func GetPosts() ([]models.Post, error) {
	var posts []models.Post

	query := "SELECT * FROM posts"

	rows, err := db.DB.Query(query)

	if err != nil {
		return []models.Post{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var nextPost models.Post
		err := rows.Scan(&nextPost.ID, &nextPost.Content, &nextPost.ImageURL, &nextPost.CreatedAt)

		if err != nil {
			return []models.Post{}, err
		}

		posts = append(posts, nextPost)
	}

	return posts, nil
}
