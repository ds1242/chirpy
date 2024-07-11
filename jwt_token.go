package main

import (
	"time"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
)

type UserClaim struct {
	jwt.RegisteredClaims
	Email fmt.Stringer
}

func CreateToken(email string)(string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaim {
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt: jwt.NewNumericDate(time.Now().UTC()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer: "chirpy",
			
		},
		Email: email,
	})
	fmt.Println(claims)
	return "string", nil
}