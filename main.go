package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nikbhavsar/go-blog-posts-api/controllers"
	"github.com/nikbhavsar/go-blog-posts-api/initalizers"
)

func init() {
	initalizers.LoadEnvVariables()
	initalizers.ConnectToDB()
}

func main() {

	router := gin.Default()
	router.POST("/posts", controllers.CreatePosts)
	router.GET("/posts", controllers.Posts)
	router.GET("/posts/:id", controllers.PostById)
	router.PUT("/posts/:id", controllers.UpdatePost)
	router.DELETE("/posts/:id", controllers.DeletePost)

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
