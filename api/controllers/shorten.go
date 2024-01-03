package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"url-shortner/api/storage"
	"url-shortner/api/utils"

	"github.com/go-playground/validator"
)

type ShortenRequest struct {
	LongURL string `json:"longURL" validate:"required"`
}

type ShortenResponse struct {
	ShortURL string `json:"shortURL"`
}



func ShortenController(w http.ResponseWriter, r *http.Request) {
	// Decode JSON payload into ShortenRequest struct
	var req ShortenRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Error decoding JSON payload", http.StatusBadRequest)
		return
	}
	
	// Use validator to check if the request is valid
	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Parse the URL
	parsedURL, urlErr := url.Parse(req.LongURL)
	if urlErr != nil || (parsedURL.Scheme == "" || parsedURL.Host == "") {
		http.Error(w, "Invalid URL format", http.StatusBadRequest)
		return
	}

	// Generate a short key
	shortKey := utils.GenerateShortKey()

	// Ensure short key doesn't already exist
	for _, exists := storage.UrlMap[shortKey]; exists; {
		shortKey = utils.GenerateShortKey()
	}

	// Store the mapping of short key to long URL in the map
	storage.UrlMap[shortKey] = req.LongURL

	// Print the map (for debugging purposes)
	fmt.Println(storage.UrlMap)
	domain := os.Getenv("DOMAIN")
	response := ShortenResponse{
		ShortURL: fmt.Sprintf("%s/%s", domain, shortKey),

	}

	// Encode the response as JSON and write it to the response writer
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
