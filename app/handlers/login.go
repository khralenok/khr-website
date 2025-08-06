package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/khralenok/khr-website/store"
)

// This function render an HTML for auth page
func ShowAuth(c *gin.Context) {
	posts, err := store.GetPosts()

	if err != nil {
		c.String(http.StatusInternalServerError, "Error loading posts")
		return
	}

	c.HTML(http.StatusOK, "base.html", posts)
}
