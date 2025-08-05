package store

import "github.com/khralenok/khr-website/db"

type Like struct {
	PostId    int  `json:"post_id"`
	UserId    int  `json:"user_id"`
	IsUnliked bool `json:"is_unliked"`
}

func AddLike(postId, userId int) (Like, error) {
	var newLike Like

	query := "INSERT INTO likes (post_id, user_id) VALUES ($1, $2) RETURNING *"

	err := db.DB.QueryRow(query, postId, userId).Scan(&newLike.PostId, &newLike.UserId, &newLike.IsUnliked)

	if err != nil {
		return Like{}, err
	}

	return newLike, nil
}

func DeleteLike(postId, userId int) error {
	query := "UPDATE likes SET is_unliked = TRUE WHERE post_id = $1 AND user_id = $2"

	_, err := db.DB.Exec(query, postId, userId)

	if err != nil {
		return err
	}

	return nil
}

func CountLikes(postId int64) int {
	var numOfLikes int

	query := "SELECT COUNT(*) FROM likes WHERE post_id = $1"

	err := db.DB.QueryRow(query, postId).Scan(&numOfLikes)

	if err != nil {
		return 0
	}

	return numOfLikes
}
