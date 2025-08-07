package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/khralenok/khr-website/store"
)

// This struct serve as temporary storage for comment data app get from the frontend
type NewCommentRequest struct {
	Content string `json:"content"`
}

// This function handle request for creating a new comment.
func CreateComment(c *gin.Context) {
	commentatorId := 1 // Must be replaced after USER implementation
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

	if err := store.AddComment(newCommentRequest.Content, postId, commentatorId); err != nil {
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

// This function render page with single post and related comments
func ShowComment(c *gin.Context) {
	commentId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": "Id parameter should be integer"})
		return
	}

	comment, err := store.GetComment(commentId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Can't load comments",
			"error":   err.Error(),
		})
		return
	}

	replies, err := store.GetReplies(commentId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Can't load replies",
			"error":   err.Error(),
		})
		return
	}

	c.HTML(http.StatusOK, "base.html", gin.H{
		"title":   "Khralenok - Comment",
		"comment": comment,
		"replies": replies,
	})
}
