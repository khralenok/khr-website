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
	commentatorId := c.GetInt("userID")
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

// This function handle request for updating a reply.
func UpdateReply(c *gin.Context) {
	userId := c.GetInt("userID")
	replyId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "Id parameter should be integer",
		})
		return
	}

	if !store.IsReplyCreator(userId, replyId) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Forbidden",
			"message": "Only author permitted to edit a reply",
		})
		return
	}

	var edits NewReplyRequest

	err = c.ShouldBindJSON(&edits)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "Can't parse input data",
		})
		return
	}

	if err := store.UpdateReply(edits.Content, replyId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not update reply",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Reply updated successfully",
		"content": edits.Content,
	})
}

// This function handle request for deleting a comment. If success, update is_deleted column value in db.
func DeleteReply(c *gin.Context) {
	userId := c.GetInt("userID")

	replyId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "Id parameter should be integer",
		})
		return
	}

	if !store.IsAdmin(userId) && !store.IsReplyCreator(userId, replyId) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Forbidden",
			"message": "Only an admin or creator are permitted to delete a reply",
		})
		return
	}

	if err := store.DeleteReply(replyId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Could not delete reply",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Post deleted successfully",
	})
}
