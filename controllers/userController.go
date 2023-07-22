package controllers

import (
	"github.com/RakibSiddiquee/golang-gin-jwt-auth-crud/errors"
	"github.com/RakibSiddiquee/golang-gin-jwt-auth-crud/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func Signup(c *gin.Context) {
	// Get the name, email and password from request
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"errors": errors.FormatValidationErrors(errs),
			})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})

	// Hash the password

	// Insert the user

	// Return success message
}
