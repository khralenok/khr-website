package store

import (
	"database/sql"
	"time"

	"github.com/khralenok/khr-website/db"
	"github.com/khralenok/khr-website/models"
)

// This function return array of posts ordered from latest to oldest
func GetPosts(userId int) ([]models.Post, error) {
	var posts []models.Post

	query := "SELECT p.* FROM posts p WHERE NOT EXISTS (SELECT 1 FROM deleted_posts d WHERE d.id = p.id) ORDER BY p.created_at DESC"

	rows, err := db.DB.Query(query)

	if err != nil {
		return []models.Post{}, err
	}

	defer rows.Close()

	for rows.Next() {
		nextPost, err := newPost(rows, userId)

		if err != nil {
			return []models.Post{}, err
		}

		posts = append(posts, nextPost)
	}

	return posts, nil
}

// This function return post with specified ID
func GetPost(postId, userId int) (models.Post, error) {
	var post models.Post

	query := "SELECT * FROM posts WHERE id=$1"

	rows, err := db.DB.Query(query, postId)

	if err != nil {
		return models.Post{}, err
	}

	defer rows.Close()

	for rows.Next() {
		post, err = newPost(rows, userId)

		if err == sql.ErrNoRows {
			break
		}

		if err != nil {
			return models.Post{}, err
		}
	}

	return post, nil
}

// This function insert new post to DB
func AddPost(content, filename string) error {
	query := "INSERT INTO posts(content, image_filename) VALUES ($1, $2)"

	_, err := db.DB.Exec(query, content, filename)

	if err != nil {
		return err
	}

	return nil
}

// This function update post with specific ID
func UpdatePost(content, filename string, postId int) error {
	if filename == "" {
		query := "UPDATE posts SET content=$1 WHERE id = $2"

		_, err := db.DB.Exec(query, content, postId)

		if err != nil {
			return err
		}

		return nil
	}

	query := "UPDATE posts SET content=$1, image_filename=$2 WHERE id = $3"

	_, err := db.DB.Exec(query, content, filename, postId)

	if err != nil {
		return err
	}

	return nil

}

// This function insert deleted post to deleted_posts table
func DeletePost(id int) error {
	query := "INSERT INTO deleted_posts(id) VALUES ($1)"

	_, err := db.DB.Exec(query, id)

	if err != nil {
		return err
	}

	return nil
}

// This function construct new Post struct from a result of database query
func newPost(row *sql.Rows, userId int) (models.Post, error) {
	var newPost models.Post
	var rawTime time.Time

	err := row.Scan(&newPost.ID, &newPost.Content, &newPost.AttachmentKey, &rawTime)

	if err != nil {
		return models.Post{}, err
	}

	newPost.NumOfComments = CountPostComments(newPost.ID) + CountPostReplies(newPost.ID)
	newPost.NumOfLikes = CountLikes(newPost.ID)

	if userId == 0 {
		newPost.IsLiked = false
	} else {
		newPost.IsLiked, err = CheckIfLikeExist(newPost.ID, userId)
	}

	if err != nil {
		return models.Post{}, err
	}

	newPost.CreatedAt = rawTime.Format("02 Jan 2006 15:04")

	return newPost, nil
}
