package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nikbhavsar/go-blog-posts-api/initalizers"
	"github.com/nikbhavsar/go-blog-posts-api/models"
)

func CreatePosts(c *gin.Context) {
	//Get Request Body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)
	//Create the Post
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initalizers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}
	//Return it
	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func Posts(c *gin.Context) {
	// Get posts
	var posts []models.Post
	initalizers.DB.Find(&posts)
	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})
}

func PostById(c *gin.Context) {
	//Get URL parameter
	id := c.Param("id")
	// Get posts
	var post []models.Post
	initalizers.DB.First(&post, id)
	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func UpdatePost(c *gin.Context) {
	//Get URL parameter
	id := c.Param("id")
	// Get Body data
	//Get Request Body
	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)

	// Find the Post
	var post []models.Post
	initalizers.DB.First(&post, id)

	//Update it
	initalizers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})
	//Send Response
	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func DeletePost(c *gin.Context) {
	//Get URL parameter
	id := c.Param("id")
	//Delete the Post
	initalizers.DB.Delete(&models.Post{}, id)
	//send Response
	c.Status(200)
}
