package database

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func (db *DB) CreateUser(password string, email string) (User, error) {
	dbStruct, err := db.loadDB()
	if err != nil {
		return User{}, err
	}
	passHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost) 
	if err != nil {
		log.Fatal(err)
	}
	
	newID := len(dbStruct.Users) + 1
	newUser := User{
		ID: 	newID,
		Password: passHash,
		Email: 	email,
	}

	dbStruct.Users[newID] = newUser
	writeErr := db.writeDB(dbStruct)
	if writeErr != nil{
		return User{}, writeErr
	}
	return newUser, nil
}