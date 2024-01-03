package api

import (
	"fmt"
	"net/http"
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
	http.Handle("/", s.router)
	fmt.Printf("Server running on %s\n", s.listenAddr)
	return http.ListenAndServe(s.listenAddr, nil)
}
