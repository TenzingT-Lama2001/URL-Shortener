package utils

import (
	"strings"
	"testing"
)

func TestGenerateShortKey(t *testing.T) {
	// Run the GenerateShortKey function multiple times and check for uniqueness
	uniqueKeys := make(map[string]struct{})
	numTests := 1000

	for i := 0; i < numTests; i++ {
		shortKey := GenerateShortKey()

		// Check if the length is as expected
		if len(shortKey) != 6 {
			t.Errorf("Generated short key length is not 6: %s", shortKey)
		}

		// Check if the characters are from the expected charset
		for _, char := range shortKey {
			if !strings.ContainsRune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789", char) {
				t.Errorf("Generated short key contains invalid character: %s", shortKey)
			}
		}

		// Check for uniqueness
		if _, exists := uniqueKeys[shortKey]; exists {
			t.Errorf("Generated short key is not unique: %s", shortKey)
		}

		uniqueKeys[shortKey] = struct{}{}
	}
}
