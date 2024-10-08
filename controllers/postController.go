package controllers

import (
	"simplecrudgo/initializers"
	"simplecrudgo/models"

	"github.com/gin-gonic/gin"
)

func PostCreate(c *gin.Context) {
	// Get data off req body
	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)
	// Create a post
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	// Get the posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	// Respond with them
	c.JSON(200, gin.H{
		"posts": posts,
	})

}

func PostShow(c *gin.Context) {
	id := c.Param("id")

	var post models.Post
	initializers.DB.First(&post, id)

	c.JSON(200, gin.H{
		"post": post,
	})

}

func PostsUpdate(c *gin.Context) {
	// Get the id off the URL
	id := c.Param("id")

	// Get the data off req body
	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)

	// Find the post were updating
	var post models.Post
	initializers.DB.First(&post, id)

	// Update it
	initializers.DB.Model(&post).Updates(models.Post{
		Body:  body.Body,
		Title: body.Title,
	})

	// Respond with it
	c.JSON(200, gin.H{
		"post": post,
	})

}

func PostsDelete(c *gin.Context) {
	// Get the id off the url
	id := c.Param("id")
	// Delete the post
	initializers.DB.Delete(&models.Post{}, id)

	// Respond
	c.Status(200)
}
