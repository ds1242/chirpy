package main

 import (
	"strings"
 )


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