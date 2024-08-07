package main

import (
	"encoding/json"
	"net/http"

	"github.com/ds1242/chirpy/helpers"
)


func (cfg *apiConfig) CreateUsersHandler(w http.ResponseWriter, r *http.Request) {
	type userParams struct {
		Password 	string `json:"password"`
		Email 		string `json:"email"`
	}

	decoder := json.NewDecoder(r.Body)
	params := userParams{}
	err := decoder.Decode(&params)

	if err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, "invalid request payload")
		return
	}

	userResponse, err := cfg.DB.CreateUser(params.Password, params.Email, cfg.JWTSecret)
	if err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	
	helpers.RespondWithJSON(w, http.StatusCreated, userResponse)
}