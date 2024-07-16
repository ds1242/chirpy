package main

import (
	"encoding/json"
	"net/http"

	"github.com/ds1242/chirpy/helpers"
)


func (cfg *apiConfig) UserRedUpgradeHandler(w http.ResponseWriter, r *http.Request) {
	type Params struct {
		Event	string `json:"event"`
		Data   	struct {
			UserID int `json:"user_id"`
		} `json:"data"`
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