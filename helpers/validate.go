package helpers

import (
	"errors"
	"strings"
)



func ValidateChirp(body string) (string, error){
	
	if len(body) > 140 {
		return "", errors.New("chirp is too long")
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


