package store

import (
	"github.com/khralenok/khr-website/db"
	"github.com/khralenok/khr-website/models"
)

// This function insert new user to database and return corresponding struct
func AddNewUser(username, pwdHash string) (models.User, error) {
	var newUser models.User

	query := "INSERT INTO users (username, pwd_hash) VALUES ($1, $2) RETURNING *"

	err := db.DB.QueryRow(query, username, pwdHash).Scan(&newUser.Id, &newUser.Username, &newUser.PwdHash, &newUser.Role, &newUser.CreatedAt)

	if err != nil {
		return models.User{}, err
	}

	return newUser, nil
}

// This function return user sruct in case if user with such username exists.
func GetUserById(id int) (models.User, error) {
	var user models.User

	query := "SELECT * FROM users WHERE id=$1"
	err := db.DB.QueryRow(query, id).Scan(&user.Id, &user.Username, &user.PwdHash, &user.Role, &user.CreatedAt)

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

// This function return user sruct in case if user with such username exists.
func GetUserByUsername(username string) (models.User, error) {
	var user models.User

	query := "SELECT * FROM users WHERE username=$1"
	err := db.DB.QueryRow(query, username).Scan(&user.Id, &user.Username, &user.PwdHash, &user.Role, &user.CreatedAt)

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

// This function return true if user role is admin.
func IsAdmin(userId int) bool {
	var role string

	query := "SELECT role FROM users WHERE id=$1"

	err := db.DB.QueryRow(query, userId).Scan(&role)

	if err != nil || role != "admin" {
		return false
	}

	return true
}

// This function return true if user is creator of the comment.
func IsCommentCreator(userId, commentId int) bool {
	var IsCreator bool

	query := "SELECT EXISTS (SELECT 1 FROM comments WHERE id = $1 AND commentator_id = $2)"

	err := db.DB.QueryRow(query, commentId, userId).Scan(&IsCreator)

	if err != nil {
		return false
	}

	return IsCreator
}

// This function return true if user is creator of the reply.
func IsReplyCreator(userId, replyId int) bool {
	var IsCreator bool

	query := "SELECT EXISTS (SELECT 1 FROM replies WHERE id = $1 AND commentator_id = $2)"

	err := db.DB.QueryRow(query, replyId, userId).Scan(&IsCreator)

	if err != nil {
		return false
	}

	return IsCreator
}
