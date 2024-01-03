package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUrlMapInitialization(t *testing.T) {
	// Arrange
	expected := make(map[string]string)

	// Act & Assert
	assert.Equal(t, expected, UrlMap, "UrlMap is not initialized as expected")
}
