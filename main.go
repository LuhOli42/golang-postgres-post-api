package main

import (
	"github.com/gin-gonic/gin"
	"luizafer.com/go-crud/controllers"
	"luizafer.com/go-crud/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.POST("/posts", controllers.PostCreate)
	r.GET("/posts", controllers.PostIndex)
	r.GET("/post/:id", controllers.PostShow)
	r.PUT("/post/:id", controllers.PostUpdate)
	r.DELETE("/post/:id", controllers.Delete)

	r.Run()
}
