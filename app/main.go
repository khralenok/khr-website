package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/khralenok/khr-website/db"
	"github.com/khralenok/khr-website/handlers"
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
	r.Static("/uploads/", "./uploads")

	// Indexed pages
	r.GET("/", handlers.ShowHome)
	//r.GET("/signin", func(ctx *gin.Context) {})
	//r.GET("/login", func(ctx *gin.Context) {})
	r.GET("/post/:id", handlers.ShowPost)
	r.GET("/comment/:id", handlers.ShowComment)

	// Not indexed pages
	r.GET("/workshop/post", func(c *gin.Context) { handlers.ShowWorkshop("post", false, c) })      //Creating workshop
	r.GET("/workshop/post/:id", func(c *gin.Context) { handlers.ShowWorkshop("post", true, c) })   //Editing workshop
	r.GET("workshop/comment", func(c *gin.Context) { handlers.ShowWorkshop("comment", false, c) }) //Create new comment
	r.GET("workshop/reply", func(c *gin.Context) { handlers.ShowWorkshop("reply", false, c) })     //Create new reply

	// Endpoints
	r.POST("/post", handlers.CreatePost)
	r.PUT("/post/:id", handlers.UpdatePost)
	r.PUT("/post/delete/:id", handlers.DeletePost)
	r.POST("/comment/:post_id", handlers.CreateComment)
	r.POST("/reply/:comment_id", handlers.CreateReply)
	r.POST("/like/:post_id", handlers.LikePost)
	r.PUT("/like/:post_id", handlers.UnlikePost)

	log.Println("Server running at http:localhost:" + port)

	if err := r.Run(":" + port); err != nil {
		log.Fatal("Server error:", err)
	}
}
