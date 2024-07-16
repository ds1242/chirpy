package main

import (
	"net/http"
	"fmt"
	"strconv"
	"github.com/ds1242/chirpy/helpers"
	"github.com/golang-jwt/jwt/v5"
)

func (cfg *apiConfig) DeleteChirpsHandler(w http.ResponseWriter, r *http.Request) {

	chirpID, err := strconv.Atoi(r.PathValue("chirpID"))

	if err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, "Invalid chirp ID")
		return
	}

	jwtToken, tokenErr := helpers.GetJWTAndStripBearer(w, *r)
	if tokenErr != nil {
		helpers.RespondWithError(w, http.StatusUnauthorized, tokenErr.Error())
		return 
	}

	claims := UserClaim{}

	token, parseTokenErr := jwt.ParseWithClaims(jwtToken, &claims, func(token *jwt.Token) (interface{}, error){
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(cfg.JWTSecret), nil
	})

	if parseTokenErr != nil {
		helpers.RespondWithError(w, http.StatusUnauthorized, parseTokenErr.Error())
		return
	}

	if !token.Valid {
		helpers.RespondWithError(w, http.StatusUnauthorized, "not authorized")
		return
	}
	authorID, authorConvErr := strconv.Atoi(claims.Subject)
	if authorConvErr != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, "Invalid author ID")
		return
	}


	deleteErr := cfg.DB.DeleteChirp(chirpID, authorID)
	if deleteErr != nil {
		helpers.RespondWithError(w, http.StatusForbidden, deleteErr.Error())
	}

	helpers.RespondWithJSON(w, http.StatusNoContent, "")
}