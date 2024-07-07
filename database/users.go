package database

import (
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func (db *DB) CreateUser(password string, email string) (User, error) {
	dbStruct, err := db.loadDB()
	if err != nil {
		return User{}, err
	}

	existingUser := SearchUserByEmail(dbStruct, email)
	if existingUser != nil {
		return User{}, errors.New("User already exists")
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


func (db *DB) UserLogin(password string, email string) (User, error) {
	dbStruct, err := db.loadDB()
	if err != nil {
		return User{}, err
	}

	existingUser := SearchUserByEmail(dbStruct, email)
	if existingUser == nil {
		return User{}, errors.New("User does not exist")
	}
	
	passErr := bcrypt.CompareHashAndPassword(existingUser.Password, []byte(password))
	if passErr != nil {
		return User{}, passErr
	}

	return *existingUser, nil
}

func SearchUserByEmail(dbStuct DBStructure, email string) *User {
	for _, user := range dbStuct.Users {
		if user.Email == email {
			return &user
		}
	}
	return nil
}