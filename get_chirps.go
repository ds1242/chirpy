package main

import (
	"net/http"
	"sort"
	"strconv"
	"github.com/ds1242/chirpy/helpers"
	"github.com/ds1242/chirpy/database"
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
		helpers.RespondWithError(w, http.StatusBadRequest, "cannot find chirp")
	}
	helpers.RespondWithJSON(w, http.StatusOK, chirpID)
}