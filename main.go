package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nikbhavsar/go-blog-posts-api/controllers"
	"github.com/nikbhavsar/go-blog-posts-api/initalizers"
	"github.com/nikbhavsar/go-blog-posts-api/middleware"
)

func init() {
	initalizers.LoadEnvVariables()
	initalizers.ConnectToDB()
}

func main() {

	router := gin.Default()
	//Post routes
	router.POST("/posts", controllers.CreatePosts)
	router.GET("/posts", controllers.Posts)
	router.GET("/posts/:id", controllers.PostById)
	router.PUT("/posts/:id", controllers.UpdatePost)
	router.DELETE("/posts/:id", controllers.DeletePost)

	//User routes
	router.POST("/signup", controllers.SignUp)
	router.POST("/login", controllers.Login)
	router.GET("/validate", middleware.RequireAuth, controllers.Validate)

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
