package main

import (
	"fmt"
	"github.com/RakibSiddiquee/golang-gin-jwt-auth-crud/api/router"
	"github.com/RakibSiddiquee/golang-gin-jwt-auth-crud/config"
	"github.com/RakibSiddiquee/golang-gin-jwt-auth-crud/db/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {
	fmt.Println("Hello auth")
	r := gin.Default()
	router.GetRoute(r)

	r.Run()
}
