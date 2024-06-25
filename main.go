package main

import (
	"net/http"
	"log"
)

// type Server struct {}

func healthzHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func main() {
	const filepathRoot = "."
	// const assetPath = "/assets"
	const port = "8080"

	mux := http.NewServeMux()
	mux.Handle("/app/*", http.StripPrefix("/app/", http.FileServer(http.Dir(filepathRoot))))
	mux.Handle("/assets", http.FileServer(http.Dir(filepathRoot)))
	mux.HandleFunc("/healthz", healthzHandler)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
	log.Fatal(srv.ListenAndServe())
}