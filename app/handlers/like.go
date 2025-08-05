package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/khralenok/khr-website/store"
)

func LikePost(c *gin.Context) {
	userId := 1 //Would be replaced after users implementation

	postId, err := strconv.Atoi(c.Param("post_id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": "Post Id parameter should be integer"})
		return
	}

	like, err := store.AddLike(postId, userId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not add this like",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Like was written",
		"post_id": like.PostId,
		"user_id": like.UserId,
	})

}

func UnlikePost(c *gin.Context) {
	userId := 1 //Would be replaced after users implementation

	postId, err := strconv.Atoi(c.Param("post_id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": "Post Id parameter should be integer"})
		return
	}

	if err := store.DeleteLike(postId, userId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not erase this like",
			"error":   err.Error(),
		})
		return
	}

	// 5. Return success/failure message
	c.JSON(http.StatusNoContent, gin.H{
		"message": "Like was erased",
	})
}
