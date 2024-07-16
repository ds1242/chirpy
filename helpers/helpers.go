package helpers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"errors"

)

func RespondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Printf("Responding with 5XX error: %s", msg)
	}
	type invalidChirp struct {
		Error string `json:"error"`
	}
	RespondWithJSON(w, code, invalidChirp{
		Error: msg,
	})
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		w.WriteHeader(500)
		return 
	}
	w.WriteHeader(code)
	w.Write(dat)
}


func GetJWTAndStripBearer(w http.ResponseWriter, r http.Request) (string, error){
	// Get the header token
	authHeader := r.Header.Get("Authorization")
	if !strings.HasPrefix(authHeader, "Bearer ") {
		return "", errors.New("not authorized")
	}

	// strip down the token from the header
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	return tokenString, nil
}
