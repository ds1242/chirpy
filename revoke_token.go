package main

import (
	"strings"
	"net/http"

	"github.com/ds1242/chirpy/helpers"
)


func (cfg *apiConfig) RevokeTokenHandler(w http.ResponseWriter, r *http.Request) {
	// Get the header token
	authHeader := r.Header.Get("Authorization")
	if !strings.HasPrefix(authHeader, "Bearer ") {
		helpers.RespondWithError(w, http.StatusUnauthorized, "not authorized")
		return
	}
	// strip down the token from the header
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	err := cfg.DB.RevokeRefreshToken(tokenString)

	if err != nil {
		helpers.RespondWithError(w, http.StatusUnauthorized, err.Error())
	}
	
	w.WriteHeader(http.StatusNoContent)
	
}