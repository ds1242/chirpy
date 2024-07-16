package main

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/ds1242/chirpy/helpers"
)

type Params struct {
	Event	string `json:"event"`
	Data   	struct {
		UserID int `json:"user_id"`
	} `json:"data"`
}

func (cfg *apiConfig) UserRedUpgradeHandler(w http.ResponseWriter, r *http.Request) {

	// Get the header token
	authHeader := r.Header.Get("Authorization")
	if !strings.HasPrefix(authHeader, "ApiKey ") {
		helpers.RespondWithError(w, http.StatusUnauthorized, "not authorized")
		return
	}

	// strip down the token from the header
	cleanKey := strings.TrimPrefix(authHeader, "ApiKey ")
	if cleanKey != cfg.PolkaKey {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)		
		return 
	}

	decoder := json.NewDecoder(r.Body)
	params := Params{}

	err := decoder.Decode(&params)
	if err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, "invalid request payload")
		return
	}

	if params.Event != "user.upgraded" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNoContent)		
		return 
	}	

	upgradeUserErr := cfg.DB.UpgradeUserToChirpyRed(params.Data.UserID)
	if upgradeUserErr != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)	
}