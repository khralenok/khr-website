package store

import (
	"time"

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
		var rawTime time.Time

		err := rows.Scan(&nextPost.ID, &nextPost.Content, &nextPost.ImageURL, &rawTime)

		if err != nil {
			return []models.Post{}, err
		}

		nextPost.CreatedAt = rawTime.Format("02 Jan 2006 15:04")

		posts = append(posts, nextPost)
	}

	return posts, nil
}
