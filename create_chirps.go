package main

import (
	"encoding/json"
	"net/http"
	"github.com/ds1242/chirpy/helpers"
)

func (cfg *apiConfig)CreateChirpHandler(w http.ResponseWriter, r *http.Request) {
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

	chirp, err := cfg.DB.CreateChirp(validatedBody)
	if err != nil {
		http.Error(w, "unable to create chirp", http.StatusBadRequest)
	}

	helpers.RespondWithJSON(w, http.StatusCreated, chirp)
}


