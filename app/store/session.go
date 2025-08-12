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

func GetSessionByToken(tokenHash []byte) (models.Session, error) {
	var session models.Session

	query := "SELECT s.user_id, u.role, s.token_hash, s.expires_at, s.revoked_at FROM sessions s JOIN users u ON u.id = s.user_id WHERE token_hash=$1"

	err := db.DB.QueryRow(query, tokenHash).Scan(&session.UserId, &session.Role, &session.TokenHash, &session.ExpiresAt, &session.RevokedAt)

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
