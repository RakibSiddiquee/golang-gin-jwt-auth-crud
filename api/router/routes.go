package router

import (
	"github.com/RakibSiddiquee/golang-gin-jwt-auth-crud/api/controllers"
	"github.com/RakibSiddiquee/golang-gin-jwt-auth-crud/api/middleware"
	"github.com/gin-gonic/gin"
)

func GetRoute(r *gin.Engine) {
	// User routes
	r.POST("/api/signup", controllers.Signup)
	r.POST("/api/login", controllers.Login)

	r.Use(middleware.RequireAuth)
	r.POST("/api/logout", controllers.Logout)
	userRouter := r.Group("/api/users")
	{
		userRouter.GET("/", controllers.GetUsers)
		userRouter.GET("/:id/edit", controllers.EditUser)
		userRouter.PUT("/:id/update", controllers.UpdateUser)
		userRouter.DELETE("/:id/delete", controllers.DeleteUser)
		userRouter.GET("/all-trash", controllers.GetTrashedUsers)
		userRouter.DELETE("/delete-permanent/:id", controllers.PermanentlyDeleteUser)
	}

	// Category routes
	catRouter := r.Group("/api/categories")
	{
		//catRouter.Use(middleware.RequireAuth)

		catRouter.GET("/", controllers.GetCategories)
		catRouter.POST("/create", controllers.CreateCategory)
		catRouter.GET("/:id/edit", controllers.EditCategory)
		catRouter.PUT("/:id/update", controllers.UpdateCategory)
		catRouter.DELETE("/:id/delete", controllers.DeleteCategory)
		catRouter.GET("/all-trash", controllers.GetTrashCategories)
		catRouter.DELETE("/delete-permanent/:id", controllers.DeleteCategoryPermanent)
	}

	// Post routes
	postRouter := r.Group("/api/posts")
	{
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
	commentRouter := r.Group("/api/posts/:id/comment")
	{
		commentRouter.POST("/store", controllers.CommentOnPost)
		commentRouter.GET("/:comment_id/edit", controllers.EditComment)
		commentRouter.PUT("/:comment_id/update", controllers.UpdateComment)
		commentRouter.DELETE("/:comment_id/delete", controllers.DeleteComment)
	}
}
