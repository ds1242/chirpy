package database

import (
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func (db *DB) CreateUser(password string, email string) (UserResponse, error) {
	dbStruct, err := db.loadDB()
	if err != nil {
		return UserResponse{}, err
	}

	existingUser := SearchUserByEmail(dbStruct, email)
	if existingUser != nil {
		return UserResponse{}, errors.New("User already exists")
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
		return UserResponse{}, writeErr
	}
	userResponse := createUserReponse(newUser)
	return userResponse, nil
}


func (db *DB) UserLogin(password string, email string) (UserResponse, error) {
	dbStruct, err := db.loadDB()
	if err != nil {
		return UserResponse{}, err
	}

	existingUser := SearchUserByEmail(dbStruct, email)
	if existingUser == nil {
		return UserResponse{}, errors.New("User does not exist")
	}

	passErr := bcrypt.CompareHashAndPassword(existingUser.Password, []byte(password))
	if passErr != nil {
		return UserResponse{}, passErr
	}

	userResponse := createUserReponse(*existingUser)
	return userResponse, nil
}

func SearchUserByEmail(dbStuct DBStructure, email string) *User {
	for _, user := range dbStuct.Users {
		if user.Email == email {
			return &user
		}
	}
	return nil
}

func createUserReponse(user User) UserResponse {
	return UserResponse{
		ID: 	user.ID,
		Email: 	user.Email,
	}
}
