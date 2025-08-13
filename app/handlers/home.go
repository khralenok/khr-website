package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/khralenok/khr-website/store"
)

// This function render an HTML for home page
func ShowHome(c *gin.Context) {
	userId := c.GetInt("userID")
	isAuth := true

	user, err := store.GetUserById(userId)

	if err != nil {
		isAuth = false
	}

	posts, err := store.GetPosts(userId)

	if err != nil {
		c.String(http.StatusInternalServerError, "Error loading posts")
		return
	}

	c.HTML(http.StatusOK, "base.html", gin.H{
		"title":    "Khralenok - Feed",
		"feed":     posts,
		"user":     user,
		"is_index": true,
		"is_auth":  isAuth,
	})
}
