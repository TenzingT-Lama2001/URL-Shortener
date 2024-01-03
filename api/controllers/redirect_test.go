package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"url-shortner/api/storage"

	"github.com/gorilla/mux"
)

func TestRedirectController(t *testing.T) {
	// Arrange
	router := mux.NewRouter()
	router.HandleFunc("/{shortCode}", RedirectController).Methods("GET")

	// Set up a test case in the storage
	testShortCode := "abc123"
	testLongURL := "https://example.com/original"
	storage.UrlMap[testShortCode] = testLongURL

	// Create a request with the short code in the path
	req, err := http.NewRequest("GET", "/"+testShortCode, nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	// Act
	router.ServeHTTP(rr, req)

	// Assert
	if status := rr.Code; status != http.StatusFound {
		t.Errorf("RedirectController returned wrong status code: got %v want %v", status, http.StatusFound)
	}

	expectedURL := testLongURL
	if rr.Header().Get("Location") != expectedURL {
		t.Errorf("RedirectController returned unexpected Location header: got %v want %v", rr.Header().Get("Location"), expectedURL)
	}
}

func TestRedirectControllerShortCodeNotFound(t *testing.T) {
	// Arrange
	router := mux.NewRouter()
	router.HandleFunc("/{shortCode}", RedirectController).Methods("GET")

	// Set up a test case in the storage
	testShortCode := "abc123"
	testLongURL := "https://example.com/original"
	storage.UrlMap[testShortCode] = testLongURL

	// Create a request with a non-existing short code in the path
	req, err := http.NewRequest("GET", "/nonexistent", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	// Act
	router.ServeHTTP(rr, req)

	// Assert
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("RedirectController returned wrong status code for non-existing short code: got %v want %v", status, http.StatusNotFound)
	}
}
