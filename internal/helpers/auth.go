package helpers

import (
	"github.com/RakibSiddiquee/golang-gin-jwt-auth-crud/api/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetAuthUser returns the authenticated user details from the Gin context
func GetAuthUser(c *gin.Context) *middleware.AuthUser {
	authUser, exists := c.Get("authUser")

	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get the user",
		})
		return nil
	}

	if user, ok := authUser.(middleware.AuthUser); ok {
		return &user
	}

	return nil
}

//func getAuthId(c *gin.Context) (uint, bool) {
//user, ok := GetAuthUser(c)
//
//if !ok {
//	return 0, false
//}

//userId, ok := user.ID.(uint)
//
//if !ok {
//	return 0, false
//}
//
//return userId, true
//}
