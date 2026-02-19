package utils

import (
	"log/slog"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitLogger(t *testing.T) {
	// Ensure InitLogger does not panic
	assert.NotPanics(t, func() {
		InitLogger()
	})

	// Ensure the default logger is set and usable
	logger := slog.Default()
	assert.NotNil(t, logger)

	// We can't easily check for the Handler type without reflection or internal access,
	// but we can ensure logging doesn't panic.
	assert.NotPanics(t, func() {
		logger.Info("Test log message")
	})
}
