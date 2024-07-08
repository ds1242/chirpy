package main

import (
	"time"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	jwt.RegisteredClaims
}

func CreateToken(username string)(string, error) {
	claims := CustomClaims {
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt: jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer: "chirpy",
			
		},
	}
	fmt.Println(claims)
	return "string", nil
}