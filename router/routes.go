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
	r.POST("/logout", controllers.Logout)

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

	// Post routes
	postRouter := r.Group("/posts")
	{
		postRouter.Use(middleware.RequireAuth)
		postRouter.GET("/", controllers.GetPosts)
		postRouter.POST("/create", controllers.CreatePost)
		postRouter.GET("/show/:id", controllers.ShowPost)
		postRouter.GET(":id", controllers.EditPost)
		postRouter.PUT("/:id", controllers.UpdatePost)
		postRouter.DELETE("/:id", controllers.DeletePost)
		postRouter.GET("/all-trash", controllers.GetTrashedPosts)
		postRouter.DELETE("/delete-permanent/:id", controllers.PermanentlyDeletePost)
	}

	// Comment routes
	commentRouter := r.Group("/posts/:id/comment")
	{
		commentRouter.Use(middleware.RequireAuth)
		commentRouter.POST("/store", controllers.CommentOnPost)
		commentRouter.GET("/:comment_id", controllers.EditComment)
		commentRouter.PUT("/:comment_id", controllers.UpdateComment)
	}
}
