package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"luizafer.com/go-crud/initializers"
	"luizafer.com/go-crud/models"
)

func PostCreate(c *gin.Context) {
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	if strings.TrimSpace(body.Title) == "" {
		c.Status(400)
		return
	}

	if strings.TrimSpace(body.Body) == "" {
		c.Status(400)
		return
	}

	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func PostIndex(c *gin.Context) {
	var posts []models.Post
	initializers.DB.Find(&posts)

	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostShow(c *gin.Context) {
	id := c.Param("id")

	var post models.Post
	initializers.DB.First(&post, id)

	c.JSON(200, gin.H{
		"posts": post,
	})
}

func PostUpdate(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	if strings.TrimSpace(body.Title) == "" {
		c.Status(400)
		return
	}

	if strings.TrimSpace(body.Body) == "" {
		c.Status(400)
		return
	}

	var post models.Post
	initializers.DB.First(&post, id)

	initializers.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body})

	c.JSON(200, gin.H{
		"posts": post,
	})
}

func Delete(c *gin.Context) {
	id := c.Param("id")

	initializers.DB.Delete(&models.Post{}, id)

	c.Status(201)
}
