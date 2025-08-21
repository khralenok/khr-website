package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/khralenok/khr-website/store"
)

// This function render an HTML for home page
func ShowHome(c *gin.Context) {
	userId := c.GetInt("userID")

	user, userErr := store.GetUserById(userId)

	posts, postErr := store.GetPosts(userId)

	if postErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": "Can't load posts",
		})
		return
	}

	if userErr == nil {
		c.HTML(http.StatusOK, "base.html", gin.H{
			"title":    "Khralenok - Feed",
			"feed":     posts,
			"user":     user,
			"is_index": true,
		})
		return
	}

	c.HTML(http.StatusOK, "base.html", gin.H{
		"title":    "Khralenok - Feed",
		"feed":     posts,
		"is_index": true,
	})
}
