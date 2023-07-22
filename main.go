package main

import (
	"fmt"
	"github.com/RakibSiddiquee/golang-gin-jwt-auth-crud/initializers"
	"github.com/RakibSiddiquee/golang-gin-jwt-auth-crud/router"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {
	fmt.Println("Hello auth")
	r := gin.Default()
	router.GetRoute(r)

	r.Run()
}
