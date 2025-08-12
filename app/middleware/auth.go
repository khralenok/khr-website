package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/khralenok/khr-website/store"
	"github.com/khralenok/khr-website/utilities"
)

// This function is middleware specific for API, that set userID to context in case if header include valid JWT token.
func AuthJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization failed"})
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token", "token": token.Valid})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Set("userID", int(claims["user_id"].(float64)))
		c.Next()
	}
}

// This function is middleware specific for SSR that check if there is an active session in DB, based on session cookies
func AuthSession() gin.HandlerFunc {
	return func(c *gin.Context) {
		raw, err := c.Cookie("sid")

		if err != nil || raw == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
			//c.Next()
			//return
		}

		tokenHash := utilities.TokenHash(raw)

		session, err := store.GetSessionByToken(tokenHash)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid session"})
			c.Abort()
			return
			//c.Next()
			//return
		}

		if time.Now().Before(session.ExpiresAt) && session.RevokedAt.Valid {
			err := store.UpdateSession(session.TokenHash)

			if err != nil {
				c.Next()
				return
			}

			c.Set("userID", session.UserId)
			c.Set("role", session.Role)
		}
		c.Next()
	}
}
