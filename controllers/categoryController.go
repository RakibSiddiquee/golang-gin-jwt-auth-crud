package controllers

import (
	"github.com/RakibSiddiquee/golang-gin-jwt-auth-crud/initializers"
	"github.com/RakibSiddiquee/golang-gin-jwt-auth-crud/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCategories(c *gin.Context) {
	// Get the categories
	var categories []models.Category

	initializers.DB.Find(&categories)

	// Return the categories
	c.JSON(http.StatusOK, gin.H{
		"categories": categories,
	})
}
