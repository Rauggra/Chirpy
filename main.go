package main

import (
	"net/http"
)

func NewServeMux() *http.ServeMux {
	mux := http.NewServeMux()
	// Add your routes here, e.g., mux.Handle("/", handler)
	// Serve static files from the current directory
	fileServer := http.FileServer(http.Dir("."))
	//picServer:="http://localhost:8080/assets/logo.png"
	mux.Handle("/", fileServer)
	return mux
}

func main() {
	mux := NewServeMux()

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

// This is a simple HTTP server that uses a custom ServeMux.
