package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/khralenok/khr-website/store"
)

func ShowHome(c *gin.Context) {
	posts, err := store.GetPosts()

	if err != nil {
		c.String(http.StatusInternalServerError, "Error loading posts")
		return
	}

	c.HTML(http.StatusOK, "base.html", gin.H{
		"title":   "Workshop - Create new post",
		"posts":   posts,
		"isIndex": true,
	})
}
