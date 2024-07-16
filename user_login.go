package main


import (
	"encoding/json"
	"net/http"
	"github.com/ds1242/chirpy/helpers"
)

func (cfg *apiConfig) UserLoginHandler(w http.ResponseWriter, r *http.Request) {
	type userParams struct {
		Password			string 	`json:"password"`
		Email 				string	`json:"email"`
		ExpiresInSeconds 	int 	`json:"expires_in_seconds,omitempty"`
	}

	defaultJWTExpiration := 60 * 60

	decoder := json.NewDecoder(r.Body)
	params := userParams{}

	err := decoder.Decode(&params)
	if err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, "invalid request payload")
	}

	expiresInSeconds := params.ExpiresInSeconds   // Assuming you've parsed this from request body
	if expiresInSeconds == 0 || expiresInSeconds > defaultJWTExpiration {
		expiresInSeconds = defaultJWTExpiration
	}

	userResponse, err := cfg.DB.UserLogin(params.Password, params.Email, expiresInSeconds, cfg.JWTSecret)
	if err != nil {
		helpers.RespondWithError(w, http.StatusUnauthorized, err.Error())
	}
	
	helpers.RespondWithJSON(w, http.StatusOK, userResponse)
}