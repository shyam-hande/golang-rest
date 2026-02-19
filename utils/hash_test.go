package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T) {
	password := "secret123"
	hash, err := HashPassword(password)

	assert.NoError(t, err)
	assert.NotEmpty(t, hash)
	assert.NotEqual(t, password, hash)
}

func TestCheckPassword(t *testing.T) {
	password := "secret123"
	hash, _ := HashPassword(password)

	// Test correct password
	valid, err := CheckPassword(hash, password)
	assert.NoError(t, err)
	assert.True(t, valid)

	// Test incorrect password
	valid, err = CheckPassword(hash, "wrongpassword")
	assert.Error(t, err) // bcrypt returns error on mismatch
	assert.False(t, valid)
}
