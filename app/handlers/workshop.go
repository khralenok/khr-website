package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/khralenok/khr-website/store"
)

// This function render an HTML for workshop pages
func ShowWorkshop(contentType string, isEditing bool, c *gin.Context) {
	switch contentType {
	case "post":
		if isEditing {
			postId, err := strconv.Atoi(c.Param("id"))

			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error":   "Bad Request",
					"message": "Id parameter should be integer",
				})
				return
			}

			content, err := store.GetPost(postId)

			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error":   "Internal Server Error",
					"message": "Can't load post",
				})
				return
			}

			c.HTML(http.StatusOK, "workshop.html", gin.H{
				"title":   "Khralenok - Edit Post",
				"isPost":  true,
				"content": content,
			})
		} else {
			c.HTML(http.StatusOK, "workshop.html", gin.H{
				"title":  "Khralenok - Create Post",
				"isPost": true,
			})
		}
	case "comment":
		if isEditing {
			postId, err := strconv.Atoi(c.Param("id"))

			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error":   "Bad Request",
					"message": "Id parameter should be integer",
				})
				return
			}

			content, err := store.GetPost(postId)

			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error":   "Internal Server Error",
					"message": "Can't load comment",
				})
				return
			}

			c.HTML(http.StatusOK, "workshop.html", gin.H{
				"title":     "Khralenok - Edit Comment",
				"isComment": true,
				"content":   content,
			})
		} else {
			c.HTML(http.StatusOK, "workshop.html", gin.H{
				"title":     "Khralenok - Create Comment",
				"isComment": true,
			})
		}
	case "reply":
	default:
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Wrong content type value",
		})
		return
	}
}
