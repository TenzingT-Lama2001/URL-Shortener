// controllers_test.go
package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"url-shortner/api/storage"

	"github.com/stretchr/testify/assert"
)

func TestShortenController(t *testing.T) {
	// Arrange
	resetStorage()

	// Create a request payload
	payload := ShortenRequest{LongURL: "https://www.example.com"}

	// Marshal the payload into JSON
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		t.Fatal(err)
	}

	// Create a request with the JSON payload
	req, err := http.NewRequest("POST", "/shorten", bytes.NewBuffer(payloadBytes))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	// Act
	ShortenController(rr, req)

	// Assert
	assert.Equal(t, http.StatusOK, rr.Code, "Status code is not OK")

	// Unmarshal the response JSON
	var response ShortenResponse
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	// Check if the short URL is not empty
	assert.NotEmpty(t, response.ShortURL, "Short URL is empty")

	// Check if the URLMap in storage has the correct mapping
	assert.Equal(t, payload.LongURL, storage.UrlMap[strings.TrimPrefix(response.ShortURL, "https://go-trim.tenzing121.com.np/")], "URL mapping is incorrect")

}

func resetStorage() {
	// Reset the URLMap in storage for each test
	storage.UrlMap = make(map[string]string)
}
