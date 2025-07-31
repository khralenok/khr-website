package store

import (
	"time"

	"github.com/khralenok/khr-website/db"
	"github.com/khralenok/khr-website/models"
)

// Return array of posts ordered from latest to oldest
func GetPosts() ([]models.Post, error) {
	var posts []models.Post

	query := "SELECT id, content, image_url, created_at FROM posts WHERE is_deleted=FALSE ORDER BY created_at DESC"

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

	query := "SELECT id, content, image_url, created_at FROM posts WHERE id=$1"

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

// Update post with specific ID
func UpdatePost(content, filename string, postId int) error {
	if filename == "" {
		query := "UPDATE posts SET content=$1 WHERE id = $2"

		_, err := db.DB.Exec(query, content, postId)

		if err != nil {
			return err
		}

		return nil
	}

	query := "UPDATE posts SET content=$1, image_url=$2 WHERE id = $3"

	_, err := db.DB.Exec(query, content, filename, postId)

	if err != nil {
		return err
	}

	return nil

}

// Delete post with specific ID
func DeletePost(postId int) error {
	query := "UPDATE posts SET is_deleted=TRUE WHERE id = $1"

	_, err := db.DB.Exec(query, postId)

	if err != nil {
		return err
	}

	return nil
}
