package main

import (
	// "fmt"
	"net/http"
	"sort"
	"strconv"

	"github.com/ds1242/chirpy/database"
	"github.com/ds1242/chirpy/helpers"
)

func (cfg *apiConfig)GetAllChirpsHandler(w http.ResponseWriter, r *http.Request) {
	dbChirps, err := cfg.DB.GetChirps()
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "unable to fetch chirps")
		return
	}

	chirps := []database.Chirp{}
	for _, dbChirp := range dbChirps {
		chirps = append(chirps, database.Chirp{
			ID:   dbChirp.ID,
			Body: dbChirp.Body,
		})
	}

	sort.Slice(chirps, func(i, j int) bool {
		return chirps[i].ID < chirps[j].ID
	})

	helpers.RespondWithJSON(w, http.StatusOK, chirps)
}

func (cfg *apiConfig)GetSingleChirpHandler(w http.ResponseWriter, r *http.Request) {
	chirpID, err := strconv.Atoi(r.PathValue("chirpID"))
	if err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, "Invalid chirp ID")
		return
	}
	chirp, err := cfg.DB.GetSingleChirp(chirpID)
	if err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, "error finding chirp")
		return
	}
	helpers.RespondWithJSON(w, http.StatusOK, database.Chirp{
		ID: 	chirp.ID,
		Body: 	chirp.Body,
	})
}

// create handler
// get id from http request
// 		convert id to an int
// pass int to a method to search the database