package main

import (
	"log"
	"net/http"
)

// type Server struct {}
type apiConfig struct {
	fileserverHits int
}



func main() {
	const filepathRoot = "."
	const port = "8080"

	apiCfg := &apiConfig{}

	mux := http.NewServeMux()
	mux.Handle("/app/*", apiCfg.middlewareMetricsInc(http.StripPrefix("/app", http.FileServer(http.Dir(filepathRoot)))))
	mux.Handle("/assets", http.FileServer(http.Dir(filepathRoot)))
	mux.HandleFunc("GET /api/healthz", healthzHandler)
	mux.HandleFunc("GET /admin/metrics", apiCfg.middlewareMetricsHandler)
	mux.HandleFunc("/api/reset", apiCfg.middlewareResetHandler)
	mux.HandleFunc("/api/chirps", validateChirp)
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
	log.Fatal(srv.ListenAndServe())
}
