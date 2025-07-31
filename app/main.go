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

	// Not indexed pages
	r.GET("/workshop/post", func(c *gin.Context) { handlers.ShowWorkshop("post", false, c) })    //Creating workshop
	r.GET("/workshop/post/:id", func(c *gin.Context) { handlers.ShowWorkshop("post", true, c) }) //Editing workshop
	//r.GET("workshop/comment", func(ctx *gin.Context) {})     //Create new comment
	//r.GET("workshop/reply", func(ctx *gin.Context) {})       //Create new reply
	//r.GET("workshop/comment/:id", func(ctx *gin.Context) {}) //Edit comment
	//r.GET("workshop/reply/:id", func(ctx *gin.Context) {})   //Edit reply

	// Endpoints
	r.POST("/post", handlers.CreatePost)
	r.PUT("/post/:id", handlers.UpdatePost)
	r.PUT("/post/delete/:id", handlers.DeletePost)
	//r.POST("/workshop/reply", func(ctx *gin.Context) {})
	//r.PUT("/workshop/post", func(ctx *gin.Context) {})
	//r.PUT("/workshop/comment", func(ctx *gin.Context) {})
	//r.PUT("/workshop/reply", func(ctx *gin.Context) {})

	log.Println("Server running at http:localhost:" + port)

	if err := r.Run(":" + port); err != nil {
		log.Fatal("Server error:", err)
	}
}
