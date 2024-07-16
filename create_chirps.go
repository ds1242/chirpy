package main

import (
	"encoding/json"
	"net/http"
	"fmt"
	
	"github.com/ds1242/chirpy/helpers"
	"github.com/golang-jwt/jwt/v5"

)

func (cfg *apiConfig)CreateChirpHandler(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body string `json:"body"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}

	err := decoder.Decode(&params)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return 
	}

	// TODO: Break this into it's own func
	// TODO: remove reused errs to avoid confusion
	jwtToken, tokenErr := helpers.GetJWTAndStripBearer(w, *r)
	if tokenErr != nil {
		helpers.RespondWithError(w, http.StatusUnauthorized, tokenErr.Error())
		return 
	}

	claims := UserClaim{}

	token, err := jwt.ParseWithClaims(jwtToken, &claims, func(token *jwt.Token) (interface{}, error){
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(cfg.JWTSecret), nil
	})

	if err != nil {
		fmt.Println("Error parsing token:", err)
		helpers.RespondWithError(w, http.StatusUnauthorized, "not authorized")
		return
	}

	if !token.Valid {
		helpers.RespondWithError(w, http.StatusUnauthorized, "not authorized")
		return
	}	
	
	validatedBody, err := helpers.ValidateChirp(params.Body)
	if err != nil {
		helpers.RespondWithError(w, 400, "unable to validate chirp body")
		return
	}

	chirp, err := cfg.DB.CreateChirp(validatedBody, claims.Subject)
	if err != nil {
		http.Error(w, "unable to create chirp", http.StatusBadRequest)
		return
	}

	helpers.RespondWithJSON(w, http.StatusCreated, chirp)
}


