package database

import (
	"time"
	// "fmt"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
)

type UserClaim struct {
	jwt.RegisteredClaims
}

func CreateToken(id int, expiresInSeconds int, jwtSecret string)(string, error) {
	
	claim := UserClaim {
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(expiresInSeconds))),
			IssuedAt: jwt.NewNumericDate(time.Now()),
			Issuer: "chirpy",
			Subject: strconv.Itoa(id),			
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return signedString, nil
}