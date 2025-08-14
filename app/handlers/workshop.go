package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/khralenok/khr-website/store"
)

// This function render an HTML for workshop pages
func ShowWorkshop(contentType string, isEditing bool, c *gin.Context) {
	userId := c.GetInt("userID")
	isAuth := true

	user, err := store.GetUserById(userId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "There is no authorized user",
		})
		return
	}

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

			content, err := store.GetPost(postId, 0) // FIX ME

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
				"user":    user,
				"is_auth": isAuth,
			})
		} else {
			c.HTML(http.StatusOK, "workshop.html", gin.H{
				"title":   "Khralenok - Create Post",
				"isPost":  true,
				"user":    user,
				"is_auth": isAuth,
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

			content, err := store.GetPost(postId, 0) // FIX ME

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
				"user":      user,
				"is_auth":   isAuth,
			})
		} else {
			c.HTML(http.StatusOK, "workshop.html", gin.H{
				"title":     "Khralenok - Create Comment",
				"isComment": true,
				"user":      user,
				"is_auth":   isAuth,
			})
		}
	case "reply":
		if isEditing {
			commentId, err := strconv.Atoi(c.Param("id"))

			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error":   "Bad Request",
					"message": "Id parameter should be integer",
				})
				return
			}

			content, err := store.GetPost(commentId, 0) // FIX ME

			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error":   "Internal Server Error",
					"message": "Can't load reply",
				})
				return
			}

			c.HTML(http.StatusOK, "workshop.html", gin.H{
				"title":   "Khralenok - Edit Reply",
				"isReply": true,
				"content": content,
				"user":    user,
				"is_auth": isAuth,
			})
		} else {
			c.HTML(http.StatusOK, "workshop.html", gin.H{
				"title":   "Khralenok - Create Reply",
				"isReply": true,
				"user":    user,
				"is_auth": isAuth,
			})
		}

	default:
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Wrong content type value",
		})
		return
	}
}
