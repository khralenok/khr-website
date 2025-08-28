package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/khralenok/khr-website/store"
	"github.com/khralenok/khr-website/utilities"
)

// This function render page with single post and related comments
func ShowPost(c *gin.Context) {
	postId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "Id parameter should be integer",
		})
		return
	}

	userId := c.GetInt("userID")

	user, userErr := store.GetUserById(userId)

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

	if userErr != nil {
		c.HTML(http.StatusOK, "base.html", gin.H{
			"title":    "Khralenok - Post",
			"post":     post,
			"comments": comments,
		})
		return
	}

	c.HTML(http.StatusOK, "base.html", gin.H{
		"title":    "Khralenok - Post",
		"user":     user,
		"post":     post,
		"comments": comments,
	})
}

// This function handle request for creating a new post.
func CreatePost(c *gin.Context) {
	userId := c.GetInt("userID")

	if !store.IsAdmin(userId) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Only an admin permitted to create a post",
		})
		return
	}

	content := c.PostForm("content")
	var attachment any
	attachmentType := c.PostForm("attachment-type")

	newPost, err := store.AddPost(content, attachmentType)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not save new post to DB",
			"error":   err.Error(),
		})
		return
	}

	switch attachmentType {
	case "image":
		c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, 15<<20)

		file, _, err := c.Request.FormFile("image")

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "file is required",
			})
			return
		}

		defer file.Close()

		proccesedImg, err := utilities.ProcessImage(file)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Image processing failed",
				"error":   err.Error(),
			})
			return
		}

		filename := utilities.GenerateImageFilename(newPost.ID, 0, "image")

		if err := utilities.SaveImage(filename, proccesedImg); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Failed to save file",
				"error":   err.Error(),
			})
			return
		}

		newAttachment, err := store.AddImageAttachment(newPost.ID, filename)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Can't record attachment into DB",
				"error":   err.Error(),
			})
			return
		}

		attachment = newAttachment
	case "carousel":
		var lastElementId int

		if err := c.Request.ParseMultipartForm(100 << 20); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		form := c.Request.MultipartForm
		images := form.File["images"]

		for i, imageHeader := range images {
			curIndex := i + 1

			image, err := imageHeader.Open()

			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": "Cannot open file",
					"error":   err.Error(),
				})
				return
			}

			defer image.Close()

			proccesedImg, err := utilities.ProcessImage(image)

			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": "Image processing failed",
					"error":   err.Error(),
				})
				return
			}

			filename := utilities.GenerateImageFilename(newPost.ID, curIndex, "carousel")

			if err := utilities.SaveImage(filename, proccesedImg); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": "Failed to save file",
					"error":   err.Error(),
				})
				return
			}

			lastElementId = curIndex
		}

		newAttachment, err := store.AddCarouselAttachment(newPost.ID, lastElementId)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Can't record attachment into DB",
				"error":   err.Error(),
			})
			return
		}

		attachment = newAttachment
	case "youtube":
		videoId := strings.TrimSpace(c.PostForm("video-id"))

		title, description, err := utilities.FetchYoutubeVideo(videoId)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "There is some problem with your video",
				"error":   err.Error(),
			})
			return
		}

		newAttachment, err := store.AddVideoAttachment(newPost.ID, videoId, title, description)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Can't record attachment into DB",
				"error":   err.Error(),
			})
			return
		}

		attachment = newAttachment
	}

	c.JSON(http.StatusOK, gin.H{
		"message":    "New post added successfully",
		"post":       newPost,
		"attachment": attachment,
	})
}

// This function handle request for updating a post.
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

	if err := store.UpdatePost(content, postId); err != nil {
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

	c.JSON(http.StatusOK, gin.H{
		"message": "Post has been deleted successfully",
	})
}
