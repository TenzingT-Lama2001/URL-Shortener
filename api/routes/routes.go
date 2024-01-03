package routes

import (
	"fmt"
	"net/http"
	"url-shortner/api/controllers"

	"github.com/gorilla/mux"
)

func NewRouter() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/", indexHandler).Methods("GET")
	router.HandleFunc("/shorten", controllers.ShortenController).Methods("POST")
	router.HandleFunc("/{shortCode}", controllers.RedirectController).Methods("GET")

	return router
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}
