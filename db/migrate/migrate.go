package main

import (
	"github.com/RakibSiddiquee/golang-gin-jwt-auth-crud/config"
	"github.com/RakibSiddiquee/golang-gin-jwt-auth-crud/db/initializers"
	models2 "github.com/RakibSiddiquee/golang-gin-jwt-auth-crud/internal/models"
	"log"
)

func init() {
	config.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {
	err := initializers.DB.Migrator().DropTable(models2.User{}, models2.Category{}, models2.Post{}, models2.Comment{})
	if err != nil {
		log.Fatal("Table dropping failed")
	}

	err = initializers.DB.AutoMigrate(models2.User{}, models2.Category{}, models2.Post{}, models2.Comment{})

	if err != nil {
		log.Fatal("Migration failed")
	}
}
