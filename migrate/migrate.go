package main

import (
	"github.com/RakibSiddiquee/golang-gin-jwt-auth-crud/initializers"
	"github.com/RakibSiddiquee/golang-gin-jwt-auth-crud/models"
	"log"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {
	err := initializers.DB.AutoMigrate(models.User{}, models.Post{})

	if err != nil {
		log.Fatal("Migration failed")
	}
}
