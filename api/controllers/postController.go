package controllers

import (
	"github.com/RakibSiddiquee/golang-gin-jwt-auth-crud/db/initializers"
	"github.com/RakibSiddiquee/golang-gin-jwt-auth-crud/internal/format-errors"
	"github.com/RakibSiddiquee/golang-gin-jwt-auth-crud/internal/helpers"
	"github.com/RakibSiddiquee/golang-gin-jwt-auth-crud/internal/models"
	"github.com/RakibSiddiquee/golang-gin-jwt-auth-crud/internal/validations"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"net/http"
)

// CreatePost creates a post
func CreatePost(c *gin.Context) {
	// Get user input from request
	var userInput struct {
		Title      string `json:"title" binding:"required,min=2,max=200"`
		Body       string `json:"body" binding:"required"`
		CategoryId uint   `json:"categoryId" binding:"required,min=1"`
	}

	if err := c.ShouldBindJSON(&userInput); err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"validations": validations.FormatValidationErrors(errs),
			})

			return
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	if !validations.IsExistValue("categories", "id", userInput.CategoryId) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"validations": map[string]interface{}{
				"CategoryId": "The category does not exist!",
			},
		})

		return
	}

	// Create a post
	authID := helpers.GetAuthUser(c).ID

	post := models.Post{
		Title:      userInput.Title,
		Body:       userInput.Body,
		CategoryID: userInput.CategoryId,
		UserID:     authID,
	}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		format_errors.InternalServerError(c)
		return
	}

	// Return the post
	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

// GetPosts gets all the post
func GetPosts(c *gin.Context) {
	// Get all the posts
	var posts []models.Post
	//var posts []struct {
	//	ID       uint   `json:"id"`
	//	Title    string `json:"title"`
	//	Body     string `json:"body"`
	//	Category struct {
	//		Name string `json:"name"`
	//		Slug string `json:"slug"`
	//	}
	//	User struct {
	//		ID   uint   `json:"id"`
	//		Name string `json:"name"`
	//	}
	//}

	result := initializers.DB.Preload("Category", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, name, slug")
	}).Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, name")
	}).Find(&posts)

	if result.Error != nil {
		format_errors.InternalServerError(c)
		return
	}

	//var PostResponse []struct {
	//	ID           uint   `json:"id"`
	//	Title        string `json:"title"`
	//	Body         string `json:"body"`
	//	CategoryName string `json:"categoryName"`
	//	CategorySlug string `json:"categorySlug"`
	//	UserId       uint   `json:"userId"`
	//	UserName     string `json:"userName"`
	//}
	//
	//initializers.DB.Table("posts").
	//	Joins("JOIN categories ON posts.category_id=categories.id").
	//	Joins("JOIN users ON posts.user_id=users.id").
	//	Select("posts.id, posts.title, posts.body, posts.user_id, categories.name as category_name, categories.slug as category_slug, users.name as user_name").Find(&PostResponse)

	// Return the posts
	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})
}

// ShowPost finds a post by ID
func ShowPost(c *gin.Context) {
	// Get the id from url
	id := c.Param("id")

	// Find the post
	var post models.Post
	result := initializers.DB.Preload("Category", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, name, slug")
	}).Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, name")
	}).Preload("Comments", func(db *gorm.DB) *gorm.DB {
		return db.Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, name")
		}).Select("id, post_id, user_id, body, created_at")
	}).First(&post, id)

	if err := result.Error; err != nil {
		format_errors.RecordNotFound(c, err)
		return
	}

	// Return the post
	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

// EditPost finds a post by ID
func EditPost(c *gin.Context) {
	// Get the id from url
	id := c.Param("id")

	// Find the post
	var post models.Post
	result := initializers.DB.First(&post, id)

	if err := result.Error; err != nil {
		format_errors.RecordNotFound(c, err)
		return
	}

	// Return the post
	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func UpdatePost(c *gin.Context) {
	// Get the id from url
	id := c.Param("id")

	// Get the data from request body
	var userInput struct {
		Title      string `json:"title" binding:"required,min=2,max=200"`
		Body       string `json:"body" binding:"required"`
		CategoryId uint   `json:"categoryId" binding:"required,min=1"`
	}

	if err := c.ShouldBindJSON(&userInput); err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"validations": validations.FormatValidationErrors(errs),
			})

			return
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Find the post by id
	var post models.Post
	result := initializers.DB.First(&post, id)

	if err := result.Error; err != nil {
		format_errors.RecordNotFound(c, err)
		return
	}

	// Prepare data to update
	authID := helpers.GetAuthUser(c).ID
	updatePost := models.Post{
		Title:      userInput.Title,
		Body:       userInput.Body,
		CategoryID: userInput.CategoryId,
		UserID:     authID,
	}

	// Update the post
	result = initializers.DB.Model(&post).Updates(&updatePost)

	if result.Error != nil {
		format_errors.InternalServerError(c)
		return
	}

	// Return the post

	c.JSON(http.StatusOK, gin.H{
		"post": updatePost,
	})
}

func DeletePost(c *gin.Context) {
	// Get the id from the url
	id := c.Param("id")
	var post models.Post

	result := initializers.DB.First(&post, id)
	if err := result.Error; err != nil {
		format_errors.RecordNotFound(c, err)
		return
	}

	// Delete the post
	initializers.DB.Delete(&post)

	// Return response
	c.JSON(http.StatusOK, gin.H{
		"message": "The post has been deleted successfully",
	})
}

func GetTrashedPosts(c *gin.Context) {
	// Get the posts
	var posts []models.Post

	result := initializers.DB.Unscoped().Find(&posts)

	if result.Error != nil {
		format_errors.InternalServerError(c)
		return
	}

	// Return the posts
	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})
}

func PermanentlyDeletePost(c *gin.Context) {
	// Get id from url
	id := c.Param("id")
	var post models.Post

	// Find the post
	if err := initializers.DB.First(&post, id).Error; err != nil {
		format_errors.RecordNotFound(c, err)
		return
	}

	// Delete the post
	initializers.DB.Unscoped().Delete(&post)

	// Return response
	c.JSON(http.StatusOK, gin.H{
		"message": "The post has been deleted permanently",
	})
}
