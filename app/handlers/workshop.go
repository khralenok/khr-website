package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/khralenok/khr-website/store"
)

func ShowWorkshop(contentType string, isEditing bool, c *gin.Context) {
	switch contentType {
	case "post":
		if isEditing {
			postId, err := strconv.Atoi(c.Param("id"))

			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": "Id parameter should be integer"})
				return
			}

			content, err := store.GetPost(postId)

			if err != nil {
				c.String(http.StatusInternalServerError, "Error loading post")
				return
			}

			c.HTML(http.StatusOK, "workshop.html", gin.H{
				"isPost":  true,
				"content": content,
			})
		} else {
			c.HTML(http.StatusOK, "workshop.html", gin.H{
				"isPost": true,
			})
		}
	case "comment":
	case "reply":
	default:
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Wrong content type value",
		})
		return
	}
}
