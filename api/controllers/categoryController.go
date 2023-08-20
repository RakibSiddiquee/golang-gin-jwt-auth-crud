package controllers

import (
	"github.com/RakibSiddiquee/golang-gin-jwt-auth-crud/db/initializers"
	"github.com/RakibSiddiquee/golang-gin-jwt-auth-crud/internal/format-errors"
	"github.com/RakibSiddiquee/golang-gin-jwt-auth-crud/internal/models"
	"github.com/RakibSiddiquee/golang-gin-jwt-auth-crud/internal/pagination"
	"github.com/RakibSiddiquee/golang-gin-jwt-auth-crud/internal/validations"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/gosimple/slug"
	"net/http"
	"strconv"
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

	pageStr := c.DefaultQuery("page", "1")
	page, _ := strconv.Atoi(pageStr)

	perPageStr := c.DefaultQuery("perPage", "5")
	perPage, _ := strconv.Atoi(perPageStr)

	result, err := pagination.Paginate(initializers.DB, page, perPage, nil, &categories)

	if err != nil {
		format_errors.InternalServerError(c)
		return
	}

	//result := initializers.DB.Find(&categories)
	//if result.Error != nil {
	//	format_errors.InternalServerError(c)
	//	return
	//}

	// Return the categories
	c.JSON(http.StatusOK, gin.H{
		"response": result,
	})
}

// EditCategory finds the category by ID
func EditCategory(c *gin.Context) {
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

	// Find the category by ID
	var category models.Category
	result := initializers.DB.First(&category, id)

	if err := result.Error; err != nil {
		format_errors.RecordNotFound(c, err)
		return
	}

	// Name unique validation
	if (category.Name != userInput.Name &&
		validations.IsUniqueValue("categories", "name", userInput.Name)) ||
		(category.Name != userInput.Name &&
			validations.IsUniqueValue("categories", "slug", slug.Make(userInput.Name))) {
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
	var category models.Category

	// Find the category
	result := initializers.DB.First(&category, id)

	if err := result.Error; err != nil {
		format_errors.RecordNotFound(c, err)
		return
	}

	// Delete the post
	initializers.DB.Delete(&category)

	// Return the response
	c.JSON(http.StatusOK, gin.H{
		"message": "The category has been deleted successfully",
	})
}

// GetTrashCategories fetch the all soft deleted categories
func GetTrashCategories(c *gin.Context) {
	// Get the categories
	var categories []models.Category

	pageStr := c.DefaultQuery("page", "1")
	page, _ := strconv.Atoi(pageStr)

	perPageStr := c.DefaultQuery("perPage", "5")
	perPage, _ := strconv.Atoi(perPageStr)

	result, err := pagination.Paginate(initializers.DB.Unscoped().Where("deleted_at IS NOT NULL"), page, perPage, nil, &categories)
	if err != nil {
		format_errors.InternalServerError(c)
		return
	}

	//result := initializers.DB.Unscoped().Where("deleted_at IS NOT NULL").Find(&categories)
	//if err := result.Error; err != nil {
	//	format_errors.InternalServerError(c)
	//	return
	//}

	// Return the categories
	c.JSON(http.StatusOK, gin.H{
		"result": result,
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
