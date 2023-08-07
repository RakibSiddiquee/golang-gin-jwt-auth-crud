package tests

import (
	"github.com/gin-gonic/gin"
	"testing"
)

func TestCreatUser(t *testing.T) {
	// Set gin to test mode
	gin.SetMode(gin.TestMode)

	// Recreate a clean database for testing
	DatabaseRefresh()
}
