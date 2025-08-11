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
func GetUserByUsername(username string) (models.User, error) {
	var user models.User

	query := "SELECT * FROM users WHERE username=$1"
	err := db.DB.QueryRow(query, username).Scan(&user.Id, &user.Username, &user.PwdHash, &user.Role, &user.CreatedAt)

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
