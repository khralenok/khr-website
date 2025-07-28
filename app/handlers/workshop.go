package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowWorkshop(c *gin.Context) {
	c.HTML(http.StatusOK, "workshop.html", gin.H{
		"title":   "Workshop - Create new post",
		"content": "workshop_content",
	})
}
