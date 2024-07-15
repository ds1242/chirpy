package database

import (
	"errors"
	"time"	
)


func (db *DB) RefreshJWTToken(refreshToken string, jwtSecret string) (string, error) {
	// Load DB
	dbStruct, err := db.loadDB()
	if err != nil {
		return "", err
	}

	existingUser, existingUserErr := SearchByRefreshToken(dbStruct, refreshToken)
	if existingUserErr != nil {
		return "", existingUserErr
	}
	defaultJWTExpiration := 60 * 60
	token, tokenErr := CreateToken(existingUser.ID, defaultJWTExpiration, jwtSecret)

	if tokenErr != nil {
		return "", tokenErr
	}

	return token, nil
}

func SearchByRefreshToken(dbStruct DBStructure, refreshToken string) (*User, error) {
	for _, user := range dbStruct.Users {
		if user.RefreshToken == refreshToken {
			if time.Now().UTC().After(user.RefreshExpiration) {
				return &User{}, errors.New("refresh token expired")
			}
			return &user, nil
		}
	}
	return &User{}, errors.New("user does not exist")
}