package controllers

import (
	"github.com/RakibSiddiquee/golang-gin-jwt-auth-crud/initializers"
	"github.com/RakibSiddiquee/golang-gin-jwt-auth-crud/models"
	"github.com/RakibSiddiquee/golang-gin-jwt-auth-crud/validations"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/gosimple/slug"
	"net/http"
)

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
	if err := initializers.DB.Where("name = ?", userInput.Name).Or("slug = ?", slug.Make(userInput.Name)).First(&models.Category{}).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{
			"validations": map[string]interface{}{
				"Name": "The name is already exist!",
			},
		})

		return
	}

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

func GetCategories(c *gin.Context) {
	// Get the categories
	var categories []models.Category

	initializers.DB.Find(&categories)

	// Return the categories
	c.JSON(http.StatusOK, gin.H{
		"categories": categories,
	})
}
