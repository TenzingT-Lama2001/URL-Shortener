package api

import (
	"fmt"
	"net/http"

	"github.com/rs/cors"
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

	c := cors.New(cors.Options{
        AllowedOrigins: []string{"http://localhost:3001"},
        AllowCredentials: true,
		AllowedMethods: []string{"GET", "POST", "DELETE", "PUT", "HEAD", "OPTIONS"},
		AllowedHeaders: []string{"*"},

		
    })

	http.Handle("/", c.Handler(s.router))
	fmt.Printf("Server running on %s\n", s.listenAddr)
	return http.ListenAndServe(s.listenAddr, nil)
}
