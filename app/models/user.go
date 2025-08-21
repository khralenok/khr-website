package models

import "time"

type User struct {
	Id             int       `json:"id"`
	Email          string    `json:"email"`
	DisplayName    string    `json:"display_name"`
	PwdHash        string    `json:"pwd_hash"`
	Role           string    `json:"role"`
	AvatarFilename string    `json:"filename"`
	CreatedAt      time.Time `json:"created_at"`
}
