package main

import (
	"log"
	"net/http"
	"os"

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
