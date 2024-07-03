package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/ds1242/chirpy/helpers"
	"github.com/ds1242/chirpy/database"
)

func CreateChirpHandler(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		type parameters struct {
			Body string `json:"body"`
		}
		decoder := json.NewDecoder(r.Body)
		params := parameters{}
		err := decoder.Decode(&params)
		if err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return 
		}
		
		validatedBody, err := helpers.ValidateChirp(params.Body)
		if err != nil {
			helpers.RespondWithError(w, 400, "unable to validate chirp body")
		}

		chirp, err := db.CreateChirp(validatedBody)
		if err != nil {
			http.Error(w, "unable to create chirp", http.StatusBadRequest)
		}

		helpers.RespondWithJSON(w, http.StatusCreated, chirp)
	}
}


func GetAllChirps(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		chirps, err := db.GetChirps()
		if err != nil {
			helpers.RespondWithError(w, 400, "unable to fetch chirps")
		}
		helpers.RespondWithJSON(w, http.StatusOK, chirps)
	}
}