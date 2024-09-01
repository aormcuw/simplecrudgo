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
