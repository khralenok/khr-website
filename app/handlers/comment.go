package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/khralenok/khr-website/store"
)

type NewCommentRequest struct {
	Content string `json:"content"`
}

// Handle request for creating a new comment.
func CreateComment(c *gin.Context) {
	postId, err := strconv.Atoi(c.Param("post_id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "Post id parameter should be integer",
		})
		return
	}

	var newCommentRequest NewCommentRequest

	if err := c.ShouldBindJSON(&newCommentRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "Wrong input json format",
		})
		return
	}

	if err := store.AddComment(newCommentRequest.Content, postId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not save new comment to DB",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"content": newCommentRequest.Content,
	})
}
