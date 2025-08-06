package store

import "github.com/khralenok/khr-website/db"

// This struct match the row from likes table and needed as storage for data returned from it
type Like struct {
	PostId    int  `json:"post_id"`
	UserId    int  `json:"user_id"`
	IsUnliked bool `json:"is_unliked"`
}

// This function add user's like to specific post or restore it if it were deleted previously
func AddLike(postId, userId int) (Like, error) {
	var newLike Like

	isDeleted, err := CheckIfLikeDeleted(postId, userId)

	if err != nil {
		return Like{}, err
	}

	query := "INSERT INTO likes (post_id, user_id) VALUES ($1, $2) RETURNING *"

	if isDeleted {
		query = "UPDATE likes SET is_unliked = FALSE WHERE post_id = $1 AND user_id = $2 RETURNING *"
	}

	err = db.DB.QueryRow(query, postId, userId).Scan(&newLike.PostId, &newLike.UserId, &newLike.IsUnliked)

	if err != nil {
		return Like{}, err
	}

	return newLike, nil
}

// This function mark user's like on specific post as deleted
func DeleteLike(postId, userId int) error {
	query := "UPDATE likes SET is_unliked = TRUE WHERE post_id = $1 AND user_id = $2"

	_, err := db.DB.Exec(query, postId, userId)

	if err != nil {
		return err
	}

	return nil
}

// This function return amount of undeleted likes on specific post
func CountLikes(postId int) int {
	var numOfLikes int

	query := "SELECT COUNT(*) FROM likes WHERE post_id = $1 AND NOT is_unliked"

	err := db.DB.QueryRow(query, postId).Scan(&numOfLikes)

	if err != nil {
		return 0
	}

	return numOfLikes
}

// This function return "true" in case when like exist or an error, if input is incorrect
func CheckIfLikeExist(postId, userId int) (bool, error) {
	var exist bool

	query := "SELECT EXISTS (SELECT 1 FROM likes WHERE post_id = $1 AND user_id = $2 AND NOT is_unliked)"

	err := db.DB.QueryRow(query, postId, userId).Scan(&exist)

	if err != nil {
		return false, err
	}

	return exist, nil
}

// This function return "true" in case when like exist but mark as deleted or an error, if input is incorrect
func CheckIfLikeDeleted(postId, userId int) (bool, error) {
	var exist bool

	query := "SELECT EXISTS (SELECT 1 FROM likes WHERE post_id = $1 AND user_id = $2 AND is_unliked)"

	err := db.DB.QueryRow(query, postId, userId).Scan(&exist)

	if err != nil {
		return false, err
	}

	return exist, nil
}
