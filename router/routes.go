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

	r.Use(middleware.RequireAuth)
	r.POST("/logout", controllers.Logout)

	// Category routes
	catRouter := r.Group("/categories")
	{
		catRouter.Use(middleware.RequireAuth)

		catRouter.GET("/", controllers.GetCategories)
		catRouter.POST("/create", controllers.CreateCategory)
		catRouter.GET("/:id/edit", controllers.EditCategory)
		catRouter.PUT("/:id/update", controllers.UpdateCategory)
		catRouter.DELETE("/:id/delete", controllers.DeleteCategory)
		catRouter.GET("/all-trash", controllers.GetTrashCategories)
		catRouter.DELETE("/delete-permanent/:id", controllers.DeleteCategoryPermanent)
	}

	// Post routes
	postRouter := r.Group("/posts")
	{
		postRouter.Use(middleware.RequireAuth)
		postRouter.GET("/", controllers.GetPosts)
		postRouter.POST("/create", controllers.CreatePost)
		postRouter.GET("/:id/show", controllers.ShowPost)
		postRouter.GET(":id/edit", controllers.EditPost)
		postRouter.PUT("/:id/update", controllers.UpdatePost)
		postRouter.DELETE("/:id/delete", controllers.DeletePost)
		postRouter.GET("/all-trash", controllers.GetTrashedPosts)
		postRouter.DELETE("/delete-permanent/:id", controllers.PermanentlyDeletePost)
	}

	// Comment routes
	commentRouter := r.Group("/posts/:id/comment")
	{
		commentRouter.Use(middleware.RequireAuth)
		commentRouter.POST("/store", controllers.CommentOnPost)
		commentRouter.GET("/:comment_id/edit", controllers.EditComment)
		commentRouter.PUT("/:comment_id/update", controllers.UpdateComment)
		commentRouter.DELETE("/:comment_id/delete", controllers.DeleteComment)
	}
}
