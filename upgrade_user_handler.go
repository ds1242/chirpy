package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ds1242/chirpy/helpers"
)


func (cfg *apiConfig) UserRedUpgradeHandler(w http.ResponseWriter, r *http.Request) {
	type DataStruct struct{
		UserID 	string `json:"user_id"`
	}
	type Params struct {
		Event	string `json:"event"`
		Data   	DataStruct 
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
	userIdInt, convErr := strconv.Atoi(params.Data.UserID)
	if convErr != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNoContent)	
		return 
	}

	upgradeUserErr := cfg.DB.UpgradeUserToChirpyRed(userIdInt)
	if upgradeUserErr != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)	
}