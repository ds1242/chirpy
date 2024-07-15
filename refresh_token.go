package main

import (
	"strings"
	"net/http"

	"github.com/ds1242/chirpy/helpers"
)


func (cfg *apiConfig) RefreshToken(w http.ResponseWriter, r *http.Request) {
	// Get the header token
	authHeader := r.Header.Get("Authorization")
	if !strings.HasPrefix(authHeader, "Bearer ") {
		helpers.RespondWithError(w, http.StatusUnauthorized, "not authorized")
		return
	}
	// strip down the token from the header
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	
	// call refresh token func
	token, err := cfg.DB.RefreshJWTToken(tokenString, cfg.JWTSecret)
	if err != nil {
		helpers.RespondWithError(w, http.StatusUnauthorized, err.Error())
	}
	// return new token if no errors
	helpers.RespondWithJSON(w, http.StatusCreated, token)

}