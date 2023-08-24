package main

import (
	"github.com/nikbhavsar/go-blog-posts-api/initalizers"
	"github.com/nikbhavsar/go-blog-posts-api/models"
)

func init() {
	initalizers.LoadEnvVariables()
	initalizers.ConnectToDB()
}

func main() {
	//	initalizers.DB.AutoMigrate(&models.Post{})
	initalizers.DB.AutoMigrate(&models.User{})

}
