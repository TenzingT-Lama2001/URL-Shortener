package controllers

import (
	"net/http"
	"url-shortner/api/storage"

	"github.com/gorilla/mux"
)

func RedirectController(w http.ResponseWriter, r *http.Request) {
	//Get the short code from the request
	vars := mux.Vars(r)
	shortCode := vars["shortCode"]

	//Get the long URL from the storage
	longURL, exists := storage.UrlMap[shortCode]

	//If the short code doesn't exist, return a 404
	if !exists {
		errorMessage := "Short code " + shortCode + " not found"
		http.Error(w, errorMessage, http.StatusNotFound)
		return
	}

	//Redirect to the long URL
	http.Redirect(w, r, longURL, http.StatusFound)
}
