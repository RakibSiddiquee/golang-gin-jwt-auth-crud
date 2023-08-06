package router

import (
	controllers2 "github.com/RakibSiddiquee/golang-gin-jwt-auth-crud/api/controllers"
	"github.com/RakibSiddiquee/golang-gin-jwt-auth-crud/api/middleware"
	"github.com/gin-gonic/gin"
)

func GetRoute(r *gin.Engine) {
	// User routes
	r.POST("/api/signup", controllers2.Signup)
	r.POST("/api/login", controllers2.Login)

	r.Use(middleware.RequireAuth)
	r.POST("/api/logout", controllers2.Logout)

	// Category routes
	catRouter := r.Group("/api/categories")
	{
		catRouter.Use(middleware.RequireAuth)

		catRouter.GET("/", controllers2.GetCategories)
		catRouter.POST("/create", controllers2.CreateCategory)
		catRouter.GET("/:id/edit", controllers2.EditCategory)
		catRouter.PUT("/:id/update", controllers2.UpdateCategory)
		catRouter.DELETE("/:id/delete", controllers2.DeleteCategory)
		catRouter.GET("/all-trash", controllers2.GetTrashCategories)
		catRouter.DELETE("/delete-permanent/:id", controllers2.DeleteCategoryPermanent)
	}

	// Post routes
	postRouter := r.Group("/api/posts")
	{
		postRouter.Use(middleware.RequireAuth)
		postRouter.GET("/", controllers2.GetPosts)
		postRouter.POST("/create", controllers2.CreatePost)
		postRouter.GET("/:id/show", controllers2.ShowPost)
		postRouter.GET(":id/edit", controllers2.EditPost)
		postRouter.PUT("/:id/update", controllers2.UpdatePost)
		postRouter.DELETE("/:id/delete", controllers2.DeletePost)
		postRouter.GET("/all-trash", controllers2.GetTrashedPosts)
		postRouter.DELETE("/delete-permanent/:id", controllers2.PermanentlyDeletePost)
	}

	// Comment routes
	commentRouter := r.Group("/api/posts/:id/comment")
	{
		commentRouter.Use(middleware.RequireAuth)
		commentRouter.POST("/store", controllers2.CommentOnPost)
		commentRouter.GET("/:comment_id/edit", controllers2.EditComment)
		commentRouter.PUT("/:comment_id/update", controllers2.UpdateComment)
		commentRouter.DELETE("/:comment_id/delete", controllers2.DeleteComment)
	}
}
