package router

import (
	"github.com/RakibSiddiquee/golang-gin-jwt-auth-crud/controllers"
	"github.com/gin-gonic/gin"
)

func GetRoute(r *gin.Engine) {
	r.POST("/signup", controllers.Signup)
}
