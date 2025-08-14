package store

import (
	"time"

	"github.com/khralenok/khr-website/db"
	"github.com/khralenok/khr-website/models"
)

// This function insert a new session into sessions table. Return error if insering fail.
func StartNewSession(userID int, tokenHash []byte, expiryDate time.Time, ip, userAgent string) error {
	query := "INSERT INTO sessions(user_id, token_hash, expires_at, ip, user_agent) VALUES($1,$2,$3,$4,$5)"

	_, err := db.DB.Exec(query, userID, tokenHash, expiryDate, ip, userAgent)

	if err != nil {
		return err
	}

	return nil
}

func RevokeSession(tokenHash []byte) error {
	query := "UPDATE sessions SET revoked_at = now() WHERE token_hash=$1"

	_, err := db.DB.Exec(query, tokenHash)

	if err != nil {
		return err
	}

	return nil
}

func GetSessionByToken(tokenHash []byte) (models.Session, error) {
	var session models.Session

	query := "SELECT user_id, token_hash, expires_at, revoked_at FROM sessions WHERE token_hash=$1"

	err := db.DB.QueryRow(query, tokenHash).Scan(&session.UserId, &session.TokenHash, &session.ExpiresAt, &session.RevokedAt)

	if err != nil {
		return models.Session{}, err
	}

	return session, nil
}

func UpdateSession(tokenHash []byte) error {
	query := "UPDATE sessions SET last_seen_at = now() WHERE token_hash=$1"

	_, err := db.DB.Exec(query, tokenHash)

	if err != nil {
		return err
	}
	return nil
}
