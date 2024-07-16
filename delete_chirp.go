package main

import (
	"encoding/json"
	"net/http"
	"fmt"
	
	"github.com/ds1242/chirpy/helpers"
	"github.com/golang-jwt/jwt/v5"
)

func (cfg *apiConfig) DeleteChirpsHandler(w http.ResponseWriter, r *http.Request) {
	
}