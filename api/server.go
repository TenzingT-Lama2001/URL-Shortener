package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"
)

type Server struct {
	listenAddr string
	router     http.Handler
}

func NewServer(listenAddr string, router http.Handler) *Server {
	return &Server{
		listenAddr: listenAddr,
		router:     router,
	}
}

func (s *Server) Start() error {

	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	http.Handle("/", corsHandler(s.router))
	http.Handle("/", s.router)
	fmt.Printf("Server running on %s\n", s.listenAddr)
	return http.ListenAndServe(s.listenAddr, nil)
}
