package main

import (
	"luizafer.com/go-crud/initializers"
	"luizafer.com/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
