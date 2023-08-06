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
	err := initializers.DB.Migrator().DropTable(models.User{}, models.Category{}, models.Post{}, models.Comment{})
	if err != nil {
		log.Fatal("Migration failed")
	}

	err = initializers.DB.AutoMigrate(models.User{}, models.Category{}, models.Post{}, models.Comment{})

	if err != nil {
		log.Fatal("Migration failed")
	}
}
