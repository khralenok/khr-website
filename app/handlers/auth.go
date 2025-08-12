package handlers

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/khralenok/khr-website/middleware"
	"github.com/khralenok/khr-website/store"
	"github.com/khralenok/khr-website/utilities"
)

// This struct needed to write in json that come from a frontend
type AuthInputs struct {
	Username string `json:"username"`
	Password string `json:"pwd"`
}

// This function render an HTML for auth page
func ShowAuth(authType string, c *gin.Context) {
	switch authType {
	case "login":
		c.HTML(http.StatusOK, "auth.html", gin.H{
			"title":   "Khralenok - Login",
			"heading": "Login to your account",
			authType:  true,
		})
	case "signin":
		c.HTML(http.StatusOK, "auth.html", gin.H{
			"title":   "Khralenok - Signin",
			"heading": "Create new account",
			authType:  true,
		})
	case "logout":
		c.HTML(http.StatusOK, "auth.html", gin.H{
			"title":  "Khralenok - Logout",
			authType: true,
		})
	default:
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Wrong authType value",
		})
		return
	}
}

// This function get login inputs create a sign about a new session in DB and return JWT Token if authenticateion is succeed.
func LoginUser(c *gin.Context) {
	var input AuthInputs

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "Invalid input format",
		})
		return
	}

	user, err := store.GetUserByUsername(input.Username)

	// 1. Examine password

	if err != nil {
		return
	}

	if !middleware.CheckPasswordHash(strings.TrimSpace(input.Password), strings.TrimSpace(user.PwdHash)) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Status Unauthorized",
			"message": "Invalid credentials",
		})
		return
	}

	// 2. Generate JWT token

	token, err := middleware.GenerateJWT(user.Id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": "Token generation failed",
		})
		return
	}

	// 3. Generate cookie-token
	rawToken, _ := utilities.NewRawToken(32)
	tokenHash := utilities.TokenHash(rawToken)

	// 4. Creating a new session
	expiryDate := time.Now().Add(7 * 24 * time.Hour)

	if err := store.StartNewSession(user.Id, tokenHash, expiryDate, c.ClientIP(), c.Request.UserAgent()); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": "Authentication failed",
		})
	}

	c.SetSameSite(http.SameSiteLaxMode)
	maxAge := int(time.Until(expiryDate).Seconds())
	c.SetCookie("sid", rawToken, maxAge, "/", "", true, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"token":   token,
	})
}

// This function add new user in database or cause an http error
func CreateUser(c *gin.Context) {
	var input AuthInputs

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "Invalid input format",
		})
	}

	pwdHash, err := middleware.HashPassword(strings.TrimSpace(input.Password))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": "Password encryption failed"},
		)
		return
	}

	user, err := store.AddNewUser(input.Username, pwdHash)

	if err != nil {
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "Created",
		"user":   user,
	})
}
