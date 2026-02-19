package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestLoggerMiddleware(t *testing.T) {
	// Set Gin to test mode to suppress debug output during tests
	gin.SetMode(gin.TestMode)

	// Create a new router with the middleware
	r := gin.New()
	r.Use(LoggerMiddleware())

	// Define a test route
	r.GET("/test-log", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	// Create a test request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test-log", nil)

	// Serve the request
	r.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)
}
