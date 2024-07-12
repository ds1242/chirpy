package main


import (
	"encoding/json"
	"net/http"
	"github.com/ds1242/chirpy/helpers"
)

func (cfg *apiConfig) UserLogin(w http.ResponseWriter, r *http.Request) {
	type userParams struct {
		Password			string 	`json:"password"`
		Email 				string	`json:"email"`
		ExpiresInSeconds 	*int 	`json:"expiresInSeconds,omitempty"`
	}

	defaultExpiration := 24 * 60 * 60

	decoder := json.NewDecoder(r.Body)
	params := userParams{}

	err := decoder.Decode(&params)
	if err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, "invalid request payload")
	}

	if params.ExpiresInSeconds == nil {
		params.ExpiresInSeconds = &defaultExpiration
	}

	if *params.ExpiresInSeconds > defaultExpiration {
		params.ExpiresInSeconds = &defaultExpiration
	}

	
	userResponse, err := cfg.DB.UserLogin(params.Password, params.Email, params.ExpiresInSeconds, cfg.JWTSecret)
	if err != nil {
		helpers.RespondWithError(w, http.StatusUnauthorized, err.Error())
	}
	
	helpers.RespondWithJSON(w, http.StatusOK, userResponse)
}