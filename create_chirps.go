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

	jwtToken, tokenErr := helpers.GetJWTAndStripBearer(w, *r)
	if tokenErr != nil {
		helpers.RespondWithError(w, http.StatusUnauthorized, tokenErr.Error())
		return 
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
		return
	}

	chirp, err := cfg.DB.CreateChirp(validatedBody, jwtToken)
	if err != nil {
		http.Error(w, "unable to create chirp", http.StatusBadRequest)
		return
	}

	helpers.RespondWithJSON(w, http.StatusCreated, chirp)
}


