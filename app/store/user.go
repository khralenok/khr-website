package store

import (
	"database/sql"

	"github.com/khralenok/khr-website/db"
	"github.com/khralenok/khr-website/models"
)

// This function insert new user to database and return corresponding struct
func AddNewUser(email, displayName, pwdHash string) (models.User, error) {
	var user models.User

	query := "INSERT INTO users (email, display_name, pwd_hash) VALUES ($1, $2, $3) RETURNING *"

	rows, err := db.DB.Query(query, email, displayName, pwdHash)

	if err != nil {
		return models.User{}, err
	}

	defer rows.Close()

	for rows.Next() {
		user, err = newUser(rows)

		if err == sql.ErrNoRows {
			break
		}

		if err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

// This function return user sruct in case if user with such username exists.
func GetUserById(id int) (models.User, error) {
	var user models.User

	query := "SELECT * FROM users WHERE id=$1"

	rows, err := db.DB.Query(query, id)

	if err != nil {
		return models.User{}, err
	}

	defer rows.Close()

	for rows.Next() {
		user, err = newUser(rows)

		if err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

// This function return user sruct in case if user with such username exists.
func GetUserByEmail(email string) (models.User, error) {
	var user models.User

	query := "SELECT * FROM users WHERE email=$1"

	rows, err := db.DB.Query(query, email)

	if err != nil {
		return models.User{}, err
	}

	defer rows.Close()

	for rows.Next() {
		user, err = newUser(rows)

		if err == sql.ErrNoRows {
			break
		}

		if err != nil {
			return models.User{}, err
		}
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

func newUser(row *sql.Rows) (models.User, error) {
	var newUser models.User

	err := row.Scan(&newUser.Id, &newUser.Email, &newUser.DisplayName, &newUser.PwdHash, &newUser.Role, &newUser.AvatarFilename, &newUser.CreatedAt)

	if err != nil {
		return models.User{}, err
	}

	return newUser, nil
}
