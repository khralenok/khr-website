package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/khralenok/khr-website/db"
	"github.com/khralenok/khr-website/handlers"
	"github.com/khralenok/khr-website/middleware"
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
	r.GET("/", middleware.AuthSession(false), handlers.ShowHome)
	r.GET("/signin", middleware.AuthSession(false), func(c *gin.Context) { handlers.ShowAuth("signin", c) })
	r.GET("/login", middleware.AuthSession(false), func(c *gin.Context) { handlers.ShowAuth("login", c) })
	r.GET("/logout", middleware.AuthSession(false), func(c *gin.Context) { handlers.ShowAuth("logout", c) })
	r.GET("/post/:id", middleware.AuthSession(false), handlers.ShowPost)
	r.GET("/comment/:id", middleware.AuthSession(false), handlers.ShowComment)

	// Not indexed pages
	r.GET("/workshop/post", middleware.AuthSession(true), func(c *gin.Context) { handlers.ShowWorkshop("post", false, c) })
	r.GET("/workshop/post/:id", middleware.AuthSession(true), func(c *gin.Context) { handlers.ShowWorkshop("post", true, c) })
	r.GET("workshop/comment", middleware.AuthSession(true), func(c *gin.Context) { handlers.ShowWorkshop("comment", false, c) })
	r.GET("workshop/reply", middleware.AuthSession(true), func(c *gin.Context) { handlers.ShowWorkshop("reply", false, c) })

	// Endpoints
	r.POST("/signin", handlers.CreateUser)
	r.POST("/login", handlers.LoginUser)
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
