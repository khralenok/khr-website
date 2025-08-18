package handlers

import (
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/khralenok/khr-website/store"
)

// This function render page with single post and related comments
func ShowPost(c *gin.Context) {
	postId, err := strconv.Atoi(c.Param("id"))
	userId := c.GetInt("userID")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": "Id parameter should be integer"})
		return
	}

	isAuth := true

	user, err := store.GetUserById(userId)

	if err != nil {
		isAuth = false
	}

	post, err := store.GetPost(postId, userId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Can't load post data",
			"error":   err.Error(),
		})
		return
	}

	comments, err := store.GetComments(postId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Can't load comments",
			"error":   err.Error(),
		})
		return
	}

	c.HTML(http.StatusOK, "base.html", gin.H{
		"title":    "Khralenok - Post",
		"user":     user,
		"post":     post,
		"comments": comments,
		"is_auth":  isAuth,
	})
}

// This function handle request for creating a new post. If success, create new row in DB and store related files on the server.
func CreatePost(c *gin.Context) {
	userId := c.GetInt("userID")

	if !store.IsAdmin(userId) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Only an admin permitted to create a post",
		})
		return
	}

	content := c.PostForm("content")
	filename := ""

	image, err := c.FormFile("image")

	if err == nil {
		savePath := filepath.Join("uploads", filepath.Base(image.Filename))

		if err := c.SaveUploadedFile(image, savePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Could not save file",
				"error":   err.Error(),
			})
			return
		}
		filename = image.Filename
	}

	if err := store.AddPost(content, filename); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not save new post to DB",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":  "New post added successfully",
		"filename": filename,
		"content":  content,
	})
}

// This function handle request for updating a post. If success, update row in DB and store related files on the server.
func UpdatePost(c *gin.Context) {
	userId := c.GetInt("userID")

	if !store.IsAdmin(userId) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Only an admin permitted to create a post",
		})
		return
	}

	postId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": "Id parameter should be integer"})
		return
	}

	content := c.PostForm("content")
	filename := ""

	image, err := c.FormFile("image")

	if err == nil {
		savePath := filepath.Join("uploads", filepath.Base(image.Filename))

		if err := c.SaveUploadedFile(image, savePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Could not save file",
				"error":   err.Error(),
			})
			return
		}
		filename = image.Filename
	}

	if err := store.UpdatePost(content, filename, postId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not save new post to DB",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Post updated successfully",
		"content": content,
	})
}

// This function handle request for deleting a post. It add records about post and all nested comments and replies to corresponding tables.
func DeletePost(c *gin.Context) {
	userId := c.GetInt("userID")

	if !store.IsAdmin(userId) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Only an admin permitted to create a post",
		})
		return
	}

	postId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": "Id parameter should be integer"})
		return
	}

	if err := store.DeletePost(postId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not delete post",
			"error":   err.Error(),
		})
		return
	}

	if err := store.DeleteComments(postId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not delete comments",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{
		"message": "Post deleted successfully",
	})
}
