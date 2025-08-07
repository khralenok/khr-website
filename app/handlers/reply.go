package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/khralenok/khr-website/store"
)

// This struct serve as temporary storage for reply data app get from the frontend
type NewReplyRequest struct {
	Content string `json:"content"`
}

// This function handle request for creating a new comment.
func CreateReply(c *gin.Context) {
	commentatorId := 1 // Must be replaced after USER implementation
	commentId, err := strconv.Atoi(c.Param("comment_id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "Post id parameter should be integer",
		})
		return
	}

	var newReplyRequest NewReplyRequest

	if err := c.ShouldBindJSON(&newReplyRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "Wrong input json format",
		})
		return
	}

	if err := store.AddReply(newReplyRequest.Content, commentId, commentatorId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not save new comment to DB",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"content": newReplyRequest.Content,
	})
}
