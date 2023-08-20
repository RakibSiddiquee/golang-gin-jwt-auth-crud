package controllers

import (
	"github.com/RakibSiddiquee/golang-gin-jwt-auth-crud/db/initializers"
	"github.com/RakibSiddiquee/golang-gin-jwt-auth-crud/internal/format-errors"
	"github.com/RakibSiddiquee/golang-gin-jwt-auth-crud/internal/models"
	"github.com/RakibSiddiquee/golang-gin-jwt-auth-crud/internal/pagination"
	"github.com/RakibSiddiquee/golang-gin-jwt-auth-crud/internal/validations"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"strconv"
	"time"
)

// Signup function is used to create a user or signup a user
func Signup(c *gin.Context) {
	// Get the name, email and password from request
	var userInput struct {
		Name     string `json:"name" binding:"required,min=2,max=50"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6"`
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

	// Email unique validation
	if validations.IsUniqueValue("users", "email", userInput.Email) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"validations": map[string]interface{}{
				"Email": "The email is already exist!",
			},
		})
		return
	}
	//if err := initializers.DB.Where("email = ?", userInput.Email).First(&models.User{}).Error; err == nil {
	//	c.JSON(http.StatusConflict, gin.H{
	//		"validations": map[string]interface{}{
	//			"Email": "The email is already exist!",
	//		},
	//	})
	//
	//	return
	//}

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

	if result.Error != nil {
		format_errors.InternalServerError(c)
		return
	}

	// Return the user
	//user.Password = ""

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

// Login function is used to log in a user
func Login(c *gin.Context) {
	// Get the email and password from the request
	var userInput struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if c.ShouldBindJSON(&userInput) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	// Find the user by email
	var user models.User
	initializers.DB.First(&user, "email = ?", userInput.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})

		return
	}

	// Compare the password with user hashed password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})

		return
	}

	// Generate a JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign in and get the complete encoded token as a string using the .env secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create token",
		})
		return
	}

	// Set expiry time and send the token back
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{})
}

// Logout function is used to log out a user
func Logout(c *gin.Context) {
	// Clear the cookie
	c.SetCookie("Authorization", "", 0, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"successMessage": "Logout successful",
	})
}

// GetUsers function is used to get users list
func GetUsers(c *gin.Context) {
	// Get all the users
	var users []models.User

	pageStr := c.DefaultQuery("page", "1")
	page, _ := strconv.Atoi(pageStr)

	perPageStr := c.DefaultQuery("perPage", "5")
	perPage, _ := strconv.Atoi(perPageStr)

	result, err := pagination.Paginate(initializers.DB, page, perPage, nil, &users)
	if err != nil {
		format_errors.InternalServerError(c)
		return
	}

	// Return the users
	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}

// EditUser function is used to find a user by id
func EditUser(c *gin.Context) {
	// Get the id from url
	id := c.Param("id")

	// Find the user
	var user models.User
	result := initializers.DB.First(&user, id)

	if err := result.Error; err != nil {
		format_errors.RecordNotFound(c, err)
		return
	}

	// Return the user
	c.JSON(http.StatusOK, gin.H{
		"result": user,
	})
}

// UpdateUser function is used to update a user
func UpdateUser(c *gin.Context) {
	// Get the id from url
	id := c.Param("id")

	// Get the name, email and password from request
	var userInput struct {
		Name  string `json:"name" binding:"required,min=2,max=50"`
		Email string `json:"email" binding:"required,email"`
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

	// Find the user by id
	var user models.User
	result := initializers.DB.First(&user, id)

	if err := result.Error; err != nil {
		format_errors.RecordNotFound(c, err)
		return
	}

	// Email unique validation
	if user.Email != userInput.Email && validations.IsUniqueValue("users", "email", userInput.Email) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"validations": map[string]interface{}{
				"Email": "The email is already exist!",
			},
		})
		return
	}

	// Prepare data to update
	updateUser := models.User{
		Name:  userInput.Name,
		Email: userInput.Email,
	}

	// Update the user
	result = initializers.DB.Model(&user).Updates(&updateUser)

	if result.Error != nil {
		format_errors.InternalServerError(c)
		return
	}

	// Return the user
	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

// DeleteUser function is used to delete a user by id
func DeleteUser(c *gin.Context) {
	// Get the id from the url
	id := c.Param("id")
	var user models.User

	result := initializers.DB.First(&user, id)
	if err := result.Error; err != nil {
		format_errors.RecordNotFound(c, err)
		return
	}

	// Delete the user
	initializers.DB.Delete(&user)

	// Return response
	c.JSON(http.StatusOK, gin.H{
		"message": "The user has been deleted successfully",
	})
}

// GetTrashedUsers function is used to get all the trashed user
func GetTrashedUsers(c *gin.Context) {
	// Get the users
	var users []models.User

	pageStr := c.DefaultQuery("page", "1")
	page, _ := strconv.Atoi(pageStr)

	perPageStr := c.DefaultQuery("perPage", "5")
	perPage, _ := strconv.Atoi(perPageStr)

	result, err := pagination.Paginate(initializers.DB.Unscoped().Where("deleted_at IS NOT NULL"), page, perPage, nil, &users)
	if err != nil {
		format_errors.InternalServerError(c)
		return
	}

	//result := initializers.DB.Unscoped().Where("deleted_at IS NOT NULL").Find(&users)
	//if err := result.Error; err != nil {
	//	format_errors.InternalServerError(c)
	//	return
	//}

	// Return the users
	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}

// PermanentlyDeleteUser function is used to delete a user permanently
func PermanentlyDeleteUser(c *gin.Context) {
	// Get the id from url
	id := c.Param("id")
	var user models.User

	// Find the user
	if err := initializers.DB.Unscoped().First(&user, id).Error; err != nil {
		format_errors.RecordNotFound(c, err)
		return
	}

	// Delete the user
	initializers.DB.Unscoped().Delete(&user)

	// Return response
	c.JSON(http.StatusOK, gin.H{
		"message": "The user has been deleted permanently",
	})
}
