package models

import (
	"database/sql"
	"time"
)

type Session struct {
	UserId    int          `json:"user_id"`
	TokenHash []byte       `json:"token_hash"`
	ExpiresAt time.Time    `json:"expires_at"`
	RevokedAt sql.NullTime `json:"revoked_at"`
}
