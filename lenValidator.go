package main

import (
	"encoding/json"
	"net/http"
)



func chirpValidator(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body string `json:"body"`
	}

	type validChirp struct {
		Cleaned_body string `json:"cleaned_body"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}
	if len(params.Body) > 140 {
		respondWithError(w, http.StatusBadRequest, "Chirp is too long")
		return 
	}

	cleanBody := chirpCleaner(params.Body)
	

	respondWithJSON(w, http.StatusOK, validChirp{
		Cleaned_body: cleanBody,
	})
}