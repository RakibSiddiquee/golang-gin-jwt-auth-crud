package controllers

import (
	format_errors "github.com/RakibSiddiquee/golang-gin-jwt-auth-crud/format-errors"
	"github.com/RakibSiddiquee/golang-gin-jwt-auth-crud/initializers"
	"github.com/RakibSiddiquee/golang-gin-jwt-auth-crud/models"
	"github.com/RakibSiddiquee/golang-gin-jwt-auth-crud/validations"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/gosimple/slug"
	"net/http"
)

// CreateCategory creates a new category
func CreateCategory(c *gin.Context) {
	// Get data from request
	var userInput struct {
		Name string `json:"name" binding:"required,min=2"`
	}

	if err := c.ShouldBindJSON(&userInput); err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"validations": validations.FormatValidationErrors(errs),
			})

			return
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	// Name unique validation
	if validations.IsUniqueValue("categories", "name", userInput.Name) ||
		validations.IsUniqueValue("categories", "slug", slug.Make(userInput.Name)) {
		c.JSON(http.StatusConflict, gin.H{
			"validations": map[string]interface{}{
				"Name": "The name is already exist!",
			},
		})

		return
	}
	//if err := initializers.DB.Where("name = ?", userInput.Name).
	//	Or("slug = ?", slug.Make(userInput.Name)).
	//	First(&models.Category{}).Error; err == nil {
	//	c.JSON(http.StatusConflict, gin.H{
	//		"validations": map[string]interface{}{
	//			"Name": "The name is already exist!",
	//		},
	//	})
	//
	//	return
	//}

	// Create the category
	category := models.Category{
		Name: userInput.Name,
	}

	result := initializers.DB.Create(&category)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Cannot create category",
		})

		return
	}

	// Return the category
	c.JSON(http.StatusOK, gin.H{
		"category": category,
	})
}

// GetCategories fetch the all categories
func GetCategories(c *gin.Context) {
	// Get the categories
	var categories []models.Category

	result := initializers.DB.Find(&categories)

	if result.Error != nil {
		format_errors.InternalServerError(c)
		return
	}

	// Return the categories
	c.JSON(http.StatusOK, gin.H{
		"categories": categories,
	})
}

// FindCategory finds the category by ID
func FindCategory(c *gin.Context) {
	// Get the id from url
	id := c.Param("id")

	// Find the post
	var category models.Category
	result := initializers.DB.First(&category, id)

	if err := result.Error; err != nil {
		format_errors.RecordNotFound(c, err)
		return
	}

	// Return the post
	c.JSON(http.StatusOK, gin.H{
		"category": category,
	})
}

// UpdateCategory updates a category
func UpdateCategory(c *gin.Context) {
	// Get the id from url
	id := c.Param("id")

	// Get the date from request
	var userInput struct {
		Name string `json:"name" binding:"required,min=2"`
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

	// Name unique validation
	if validations.IsUniqueValue("categories", "name", userInput.Name) ||
		validations.IsUniqueValue("categories", "slug", slug.Make(userInput.Name)) {
		c.JSON(http.StatusConflict, gin.H{
			"validations": map[string]interface{}{
				"Name": "The name is already exist!",
			},
		})

		return
	}
	//if err := initializers.DB.Where("name = ?", userInput.Name).
	//	Or("slug = ?", slug.Make(userInput.Name)).
	//	First(&models.Category{}).Error; err == nil {
	//	c.JSON(http.StatusConflict, gin.H{
	//		"validations": map[string]interface{}{
	//			"Name": "The name is already exist!",
	//		},
	//	})
	//
	//	return
	//}

	// Find the category by ID
	var category models.Category
	result := initializers.DB.First(&category, id)

	if err := result.Error; err != nil {
		format_errors.RecordNotFound(c, err)
		return
	}

	updateCategory := models.Category{
		Name: userInput.Name,
		Slug: slug.Make(userInput.Name),
	}

	// Update the category record
	result = initializers.DB.Model(&category).Updates(updateCategory)
	if err := result.Error; err != nil {
		format_errors.InternalServerError(c)
		return
	}

	// Return the category
	c.JSON(http.StatusOK, gin.H{
		"category": updateCategory,
	})
}

// DeleteCategory deletes a category by id
func DeleteCategory(c *gin.Context) {
	// Get the id from request
	id := c.Param("id")

	// Delete the post
	result := initializers.DB.Delete(&models.Category{}, id)

	if err := result.Error; err != nil {
		format_errors.RecordNotFound(c, err)
		return
	}

	// Return the response
	c.JSON(http.StatusOK, gin.H{
		"message": "The category has been deleted successfully",
	})
}

// GetTrashCategories fetch the all soft deleted categories
func GetTrashCategories(c *gin.Context) {
	// Get the categories
	var categories []models.Category

	result := initializers.DB.Unscoped().Find(&categories)
	if err := result.Error; err != nil {
		format_errors.InternalServerError(c)
		return
	}

	// Return the categories
	c.JSON(http.StatusOK, gin.H{
		"categories": categories,
	})
}

func DeleteCategoryPermanent(c *gin.Context) {
	// Get the id from request
	id := c.Param("id")

	// Delete the post
	result := initializers.DB.Unscoped().Delete(&models.Category{}, id)

	if err := result.Error; err != nil {
		format_errors.RecordNotFound(c, err)
		return
	}

	// Return the response
	c.JSON(http.StatusOK, gin.H{
		"message": "The category has been deleted permanently",
	})
}
