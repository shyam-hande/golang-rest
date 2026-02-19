package middleware

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"rest-api/utils"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAuthenticate(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Setup Router
	r := gin.New()
	r.Use(Authenticate)
	r.GET("/protected", func(c *gin.Context) {
		userId, _ := c.Get("userId")
		c.JSON(http.StatusOK, gin.H{"userId": userId})
	})

	// Case 1: No Token
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/protected", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)

	// Case 2: Invalid Token
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/protected", nil)
	req.Header.Set("Authorization", "invalid-token")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)

	// Case 3: Valid Token
	userId := int64(999)
	token, _ := utils.GenerateToken("test@auth.com", userId)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/protected", nil)
	req.Header.Set("Authorization", token)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, float64(userId), response["userId"]) // JSON numbers are floats
}
