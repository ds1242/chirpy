package helpers

import (
	// "encoding/json"
	// "net/http"
	"errors"
	"strings"
)



func ValidateChirp(body string) (string, error){
	

	// type validChirp struct {
	// 	Cleaned_body string `json:"cleaned_body"`
	// }

	// decoder := json.NewDecoder(r.Body)
	// params := parameters{}
	// err := decoder.Decode(&params)
	// if err != nil {
	// 	respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
	// 	return false, "Couldn't decode parameters"
	// }
	if len(body) > 140 {
		// respondWithError(w, http.StatusBadRequest, "Chirp is too long")
		return "", errors.New("Chirp is too long")
	}

	badWords := map[string]struct{}{
		"kerfuffle": {},
		"sharbert":  {},
		"fornax":    {},
	}

	cleanBody := chirpCleaner(body, badWords)
	
	return cleanBody, nil 
}



func chirpCleaner(body string, badWords map[string]struct{}) string {
	bodySplit := strings.Split(body, " ")
	for i, word := range bodySplit {
		loweredWord := strings.ToLower(word)
		if _, ok := badWords[loweredWord]; ok {
			bodySplit[i] = "****"
		}
	}
	cleaned := strings.Join(bodySplit, " ")
	return cleaned
}


