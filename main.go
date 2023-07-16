package main

import (
	"fmt"
	"github.com/RakibSiddiquee/golang-gin-jwt-auth-crud/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {
	fmt.Println("Hello auth")
}
