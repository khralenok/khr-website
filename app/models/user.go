package models

import "time"

type User struct {
	Id        int       `json:"id"`
	Username  string    `json:"username"`
	PwdHash   string    `json:"pwd_hash"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}
