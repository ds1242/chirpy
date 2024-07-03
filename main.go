package main

import (
	"log"
	"net/http"
	"github.com/ds1242/chirpy/database"
	"github.com/ds1242/chirpy/handlers"
)

// type Server struct {}
type apiConfig struct {
	fileserverHits int
}



func main() {
	const filepathRoot = "."
	const port = "8080"

	apiCfg := &apiConfig{}
	
	db, err := database.NewDB("./database.json")
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	mux := http.NewServeMux()
	mux.Handle("/app/*", apiCfg.middlewareMetricsInc(http.StripPrefix("/app", http.FileServer(http.Dir(filepathRoot)))))
	mux.Handle("/assets", http.FileServer(http.Dir(filepathRoot)))
	mux.HandleFunc("GET /api/healthz", healthzHandler)
	mux.HandleFunc("GET /admin/metrics", apiCfg.middlewareMetricsHandler)
	mux.HandleFunc("/api/reset", apiCfg.middlewareResetHandler)
	mux.HandleFunc("POST /api/chirps", handlers.CreateChirpHandler(db))
	mux.HandleFunc("GET /api/chirps", handlers.GetAllChirps(db))
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
	log.Fatal(srv.ListenAndServe())
}
