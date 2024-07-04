package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ds1242/chirpy/helpers"
)


func (cfg *apiConfig) CreateUsersHandler(w http.ResponseWriter, r *http.Request) {
	type userParams struct {
		Email string `json:"email"`
	}

	decoder := json.NewDecoder(r.Body)
	params := userParams{}
	err := decoder.Decode(&params)

	if err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, "invalid request payload")
	}
	fmt.Println(params.Email)
	user, err := cfg.DB.CreateUser(params.Email)
	if err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, "unable to create a user")
	}
	helpers.RespondWithJSON(w, http.StatusCreated, user)
}