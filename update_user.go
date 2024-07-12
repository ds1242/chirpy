package main

import (
	// "encoding/json"
	"fmt"
	"net/http"
	"strings"
	"encoding/json"

	"github.com/golang-jwt/jwt/v5"
	"github.com/ds1242/chirpy/helpers"
)
type UserClaim struct {
	jwt.RegisteredClaims
}


func (cfg *apiConfig) UpdateUser(w http.ResponseWriter, r *http.Request) {

	type userParams struct {
		Password			string 	`json:"password,omitempty"`
		Email 				string	`json:"email,omitempty"`
	}

	decoder := json.NewDecoder(r.Body)
	params := userParams{}

	err := decoder.Decode(&params)
	if err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, "invalid request payload")
		return
	}

	authHeader := r.Header.Get("Authorization")
	if !strings.HasPrefix(authHeader, "Bearer ") {
		helpers.RespondWithError(w, http.StatusUnauthorized, "not authorized")
		return
	}
	
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	claims := UserClaim{}

	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error){
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

	// fmt.Println(token)
	userResponse, err := cfg.DB.UserUpdate(claims.Subject, params.Email, params.Password, tokenString)
	// fmt.Println(userResponse)
	if err != nil {
		helpers.RespondWithError(w, http.StatusUnauthorized, err.Error())
		return
	}

	helpers.RespondWithJSON(w, http.StatusOK, userResponse)
}