package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/khralenok/khr-website/db"
	"github.com/khralenok/khr-website/store"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	if err := db.Connect(); err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	defer db.DB.Close()

	r := gin.Default()

	r.LoadHTMLGlob("templates/*.html")
	r.Static("/static/", "./static")

	r.GET("/", showHome)

	r.POST("/create-post", createPost)

	log.Println("Server running at http:localhost:" + port)

	if err := r.Run(":" + port); err != nil {
		log.Fatal("Server error:", err)
	}
}

func showHome(c *gin.Context) {
	posts, err := store.GetPosts()

	if err != nil {
		c.String(http.StatusInternalServerError, "Error loading posts")
		return
	}

	c.HTML(http.StatusOK, "base.html", posts)
}

func createPost(c *gin.Context) {
	// 1. Get input
	content := c.PostForm("content")

	image, err := c.FormFile("image")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Image upload failed", "error": err.Error()})
		return
	}

	// 2. Validate input (TO DO)

	// 3. Store file in img folder
	savePath := filepath.Join("static", "img", filepath.Base(image.Filename))

	if err := c.SaveUploadedFile(image, savePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not save file",
			"error":   err.Error(),
		})
		return
	}

	// 4. Store content and link in db
	imageURL := "img/" + image.Filename
	if err := store.AddPost(content, imageURL); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not save new post to DB",
			"error":   err.Error(),
		})
		return
	}

	// 5. Return success/failure message
	c.JSON(http.StatusOK, gin.H{
		"message":  "File uploaded successfully",
		"filename": imageURL,
		"content":  content,
	})
}
