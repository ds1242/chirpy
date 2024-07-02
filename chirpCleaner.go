package main

 import (
	"strings"
 )


func chirpCleaner(body string) string {
	bodySplit := strings.Split(body, " ")
	var outputSlice []string
	for _, word := range(bodySplit) {
		if strings.ToLower(word) == "kerfuffle" || strings.ToLower(word) == "sharbert" || strings.ToLower(word) == "fornax" {
			outputSlice = append(outputSlice, "****")
			continue
		}
		outputSlice = append(outputSlice, word)
	}
	joinedBody := strings.Join(outputSlice, " ")
	return joinedBody
}