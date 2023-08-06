package format_errors

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func RecordNotFound(c *gin.Context, err error, errMessage ...string) {
	errorMessage := "The record not found"
	if len(errMessage) > 0 {
		errorMessage = errMessage[0]
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": errorMessage,
		})
		return
	}

	// Else show internal server error
	InternalServerError(c)
}

func InternalServerError(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"error": "Internal server error",
	})
	return
}
