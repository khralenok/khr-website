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
	commentatorId := c.GetInt("userID")
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

// This function handle request for updating a post. If success, update row in DB and store related files on the server.
func UpdateComment(c *gin.Context) {
	userId := c.GetInt("userID")
	commentId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "Id parameter should be integer",
		})
		return
	}

	if !store.IsCommentCreator(userId, commentId) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Only author permitted to edit a comment",
		})
		return
	}

	content := c.PostForm("content")

	if err := store.UpdateComment(content, commentId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not save new post to DB",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Comment updated successfully",
		"content": content,
	})
}

// This function handle request for deleting a comment. If success, update is_deleted column value in db.
func DeleteComment(c *gin.Context) {
	userId := c.GetInt("userID")

	commentId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": "Id parameter should be integer"})
		return
	}

	if !store.IsAdmin(userId) && !store.IsCommentCreator(userId, commentId) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Only an admin or creator are permitted to delete a comment",
		})
		return
	}

	if err := store.DeleteComment(commentId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not delete comment",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{
		"message": "Post deleted successfully",
	})
}

// This function render page with single post and related comments
func ShowComment(c *gin.Context) {
	userId := c.GetInt("userID")
	commentId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": "Id parameter should be integer"})
		return
	}

	isAuth := true

	user, err := store.GetUserById(userId)

	if err != nil {
		isAuth = false
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
		"user":    user,
		"is_auth": isAuth,
	})
}
