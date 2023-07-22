package controllers

import (
	"github.com/RakibSiddiquee/golang-gin-jwt-auth-crud/initializers"
	"github.com/RakibSiddiquee/golang-gin-jwt-auth-crud/models"
	"github.com/RakibSiddiquee/golang-gin-jwt-auth-crud/validations"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

// Signup is used to create a user or signup a user
func Signup(c *gin.Context) {
	// Get the name, email and password from request
	var userInput struct {
		Name     string `json:"name" binding:"required,min=2,max=50"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6"`
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

	// Email unique validation
	if err := initializers.DB.Where("email = ?", userInput.Email).First(&models.User{}).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{
			"error": "The email is already exist!",
		})

		return
	}

	// Hash the password
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(userInput.Password), 10)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to hash password",
		})

		return
	}

	user := models.User{
		Name:     userInput.Name,
		Email:    userInput.Email,
		Password: string(hashPassword),
	}

	// Create the user
	result := initializers.DB.Create(&user)

	log.Println(result.Error)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create user",
		})

		return
	}

	// Return the user
	//user.Password = ""

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func Login(c *gin.Context) {
	// Get the email and password from the request
	//var userInput struct {
	//	Email    string `json:"email" binding:"required,email"`
	//	Password string
	//}

	// Find the user by email

	// Compare the password with user hashed password

	// Generate a JWT token

	// Sign in and get the complete encoded token as a string using the .env secret

	// Set expiry time and send the token back
}
