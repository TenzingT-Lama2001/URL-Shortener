package api

import (
	"net/http"
	"testing"
)

func TestServer_Start(t *testing.T) {
	// Create a new router
	testRouter := http.NewServeMux()
	// Create a new server with a dynamic port
	server := NewServer(":0", testRouter)

	// Start the server in a goroutine
	go func() {
		err := server.Start()
		if err != nil {
			t.Errorf("Server failed to start: %v", err)
		}
	}()

}
