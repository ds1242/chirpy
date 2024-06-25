package main

import (
	"net/http"
	"log"
)

// type Server struct {}

// func (Server) ServeHTTP(http.ResponseWriter, *http.Request) {}

func main() {
	const filepathRoot = "."
	// const assetPath = "/assets"
	const port = "8080"

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(filepathRoot)))
	mux.Handle("/assets", http.FileServer(http.Dir(filepathRoot)))

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
	log.Fatal(srv.ListenAndServe())
}