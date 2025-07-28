package store

import (
	"time"

	"github.com/khralenok/khr-website/db"
	"github.com/khralenok/khr-website/models"
)

// Return array of posts ordered from latest to oldest
func GetPosts() ([]models.Post, error) {
	var posts []models.Post

	query := "SELECT * FROM posts ORDER BY created_at DESC"

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

// Return post with specified ID
func GetPost(postId int) (models.Post, error) {
	var post models.Post
	var rawTime time.Time

	query := "SELECT * FROM posts WHERE id=$1"

	err := db.DB.QueryRow(query, postId).Scan(&post.ID, &post.Content, &post.ImageURL, &rawTime)

	if err != nil {
		return models.Post{}, err
	}

	post.CreatedAt = rawTime.Format("02 Jan 2006 15:04")

	return post, nil
}

// Insert new post to DB
func AddPost(content, filename string) error {
	query := "INSERT INTO posts(content, image_url) VALUES ($1, $2)"

	_, err := db.DB.Exec(query, content, filename)

	if err != nil {
		return err
	}

	return nil
}
