package database

import (
	"time"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
)

type UserClaim struct {
	jwt.RegisteredClaims
}

func CreateToken(id int, expiresInSeconds int, jwtSecret string)(string, error) {
	currentTime := time.Now().UTC()
	expirationTime := time.Now().Add(time.Duration(expiresInSeconds) * time.Second).UTC()
	// create the JWT claims, which include the user ID and expiration time
	claim := UserClaim {
		jwt.RegisteredClaims{
			Issuer: "chirpy",
			IssuedAt: jwt.NewNumericDate(currentTime),
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			Subject: strconv.Itoa(id),			
		},
	}
	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	// Sign the token withthe secret
	signedString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return signedString, nil
}