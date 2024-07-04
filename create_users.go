package main

import (
	"encoding/json"
	"net/http"

	"github.com/ds1242/chirpy/helpers"
)


func (cfg *apiConfig) CreateUsersHandler(w http.ResponseWriter, r *http.Request) {
	type userParams struct {
		Email string `json:"email"`
	}

	decoder := json.Decoder(r.Body)
	params := userParams{}
	err := decoder.Decode(&params)

	if err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, "invalid request payload")
	}
}