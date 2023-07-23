package router

import (
	"github.com/RakibSiddiquee/golang-gin-jwt-auth-crud/controllers"
	"github.com/RakibSiddiquee/golang-gin-jwt-auth-crud/middleware"
	"github.com/gin-gonic/gin"
)

func GetRoute(r *gin.Engine) {
	// User routes
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)

	// Category routes
	r.GET("/categories", middleware.RequireAuth, controllers.GetCategories)
}
