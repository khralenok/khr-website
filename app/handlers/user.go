package handlers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/khralenok/khr-website/middleware"
	"github.com/khralenok/khr-website/store"
)

// This struct needed to write in json that come from a frontend
type AuthInputs struct {
	Username string `json:"username"`
	Password string `json:"pwd"`
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

// This function get login inputs and return JWT Token if authenticateion is succeed.
func LoginUser(context *gin.Context) {
	var input AuthInputs

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": "Invalid input format"})
		return
	}

	user, err := store.GetUserByUsername(input.Username)

	if err != nil {
		return
	}

	if !middleware.CheckPasswordHash(strings.TrimSpace(input.Password), strings.TrimSpace(user.PwdHash)) {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Status Unauthorized", "message": "Invalid credentials"})
		return
	}

	token, err := middleware.GenerateJWT(user.Id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error", "message": "Token generation failed"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Success", "token": token})
}
