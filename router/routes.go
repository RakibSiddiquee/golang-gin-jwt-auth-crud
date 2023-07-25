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
	catRouter := r.Group("/categories")
	{
		catRouter.Use(middleware.RequireAuth)

		catRouter.GET("/", controllers.GetCategories)
		catRouter.POST("/create", controllers.CreateCategory)
		catRouter.GET("/:id", controllers.FindCategory)
		catRouter.PUT("/:id", controllers.UpdateCategory)
		catRouter.DELETE("/:id", controllers.DeleteCategory)
		catRouter.GET("/all-trash", controllers.GetTrashCategories)
		catRouter.DELETE("/delete-permanent/:id", controllers.DeleteCategoryPermanent)
	}

}
