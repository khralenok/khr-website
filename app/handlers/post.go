package handlers

import (
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/khralenok/khr-website/store"
)

// Render page with single post and related comments
func ShowPost(c *gin.Context) {
	postId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": "Id parameter should be integer"})
		return
	}

	post, err := store.GetPost(postId)

	if err != nil {
		c.String(http.StatusInternalServerError, "Error loading post")
		return
	}

	c.HTML(http.StatusOK, "base.html", gin.H{
		"title":   "Khralenok Blog",
		"post":    post,
		"isIndex": false,
	})
}

// Handle request for creating a new post. If success, create new row in DB and store related files on the server.
func CreatePost(c *gin.Context) {
	content := c.PostForm("content")
	image, err := c.FormFile("image")
	filename := ""

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

	// 5. Return success/failure message
	c.JSON(http.StatusOK, gin.H{
		"message":  "File uploaded successfully",
		"filename": filename,
		"content":  content,
	})
}
