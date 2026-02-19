package utils

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

func TestGenerateToken(t *testing.T) {
	email := "test@example.com"
	userId := int64(101)

	token, err := GenerateToken(email, userId)

	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestVerifyToken(t *testing.T) {
	email := "test@example.com"
	userId := int64(102)

	token, _ := GenerateToken(email, userId)

	// Test valid token
	parsedUserId, err := VerifyToken(token)
	assert.NoError(t, err)
	assert.Equal(t, userId, parsedUserId)

	// Test invalid token
	_, err = VerifyToken("invalid-token-string")
	assert.Error(t, err)
}

func TestExpiredToken(t *testing.T) {
	// Manually create an expired token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  "test@example.com",
		"userId": float64(103),
		"exp":    time.Now().Add(-time.Hour).Unix(), // Expired 1 hour ago
	})
	signedToken, _ := token.SignedString([]byte(secretKey))

	_, err := VerifyToken(signedToken)
	assert.Error(t, err)
}
