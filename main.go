package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"github.com/ds1242/chirpy/database"
)

// type Server struct {}
type apiConfig struct {
	fileserverHits 	int
	DB 				*database.DB
}



func main() {
	const filepathRoot = "."
	const port = "8080"

	dbg := flag.Bool("debug", false, "Enable debug mode")
	flag.Parse()
	
	if *dbg {
		err := os.Remove("./database.json")
		if err != nil {
			log.Fatal(err)
		}
	}
	db, err := database.NewDB("./database.json")
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	
	apiCfg := &apiConfig{
		fileserverHits: 0,
		DB:				db,
	}
	mux := http.NewServeMux()
	mux.Handle("/app/*", apiCfg.middlewareMetricsInc(http.StripPrefix("/app", http.FileServer(http.Dir(filepathRoot)))))
	mux.Handle("/assets", http.FileServer(http.Dir(filepathRoot)))
	mux.HandleFunc("GET /api/healthz", healthzHandler)
	mux.HandleFunc("GET /admin/metrics", apiCfg.middlewareMetricsHandler)
	mux.HandleFunc("/api/reset", apiCfg.middlewareResetHandler)
	mux.HandleFunc("POST /api/chirps", apiCfg.CreateChirpHandler)
	mux.HandleFunc("GET /api/chirps", apiCfg.GetAllChirpsHandler)
	mux.HandleFunc("GET /api/chirps/{chirpID}", apiCfg.GetSingleChirpHandler)
	mux.HandleFunc("POST /api/users", apiCfg.CreateUsersHandler)
	mux.HandleFunc("POST /api/login", apiCfg.UserLogin)
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
	log.Fatal(srv.ListenAndServe())
}
