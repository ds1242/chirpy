package database

import (
	"errors"
	"log"	
	"strconv"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)



func (db *DB) CreateUser(password string, email string, jwtSecret string) (UserResponse, error) {
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

	expiresInSeconds := 24 * 60 * 60
	token, tokenErr := CreateToken(newUser.ID, expiresInSeconds, jwtSecret)

	if tokenErr != nil {
		return UserResponse{}, tokenErr
	}

	userResponse := createUserReponse(newUser, token)
	return userResponse, nil
}


func (db *DB) UserLogin(password string, email string, expiresInSeconds *int, jwtSecret string) (UserResponse, error) {
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

	token, tokenErr := CreateToken(existingUser.ID, *expiresInSeconds, jwtSecret)
	if tokenErr != nil {
		return UserResponse{}, tokenErr
	}

	userResponse := createUserReponse(*existingUser, token)
	return userResponse, nil
}


type UpdateUserParams struct {
    Email    string `json:"email,omitempty"`
    Password string `json:"password,omitempty"`
}

func (db *DB) UserUpdate(userID string, email string, password string, jwtSecret string) (UserResponse, error) {
	id, err := strconv.Atoi(userID)
	fmt.Println(email)
	fmt.Println(password)
	if err != nil {
		return UserResponse{}, err
	}


	dbStruct, err := db.loadDB()
	
	if err != nil {
		return UserResponse{}, err
	}
	user, userErr := GetUserByID(dbStruct, id)
	if userErr != nil {
		return UserResponse{}, userErr
	}
	
	if len(email) > 0 {
		user.Email = email
	}

	if len(password) > 0 {
		passHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost) 
		if err != nil {
			log.Fatal(err)
		}
		user.Password = passHash
	}

	dbStruct.Users[id] = *user
	
	writeErr := db.writeDB(dbStruct)
	if writeErr != nil{
		return UserResponse{}, writeErr
	}

	userResponse := createUserReponse(*user, jwtSecret)
	return userResponse, nil
}

func GetUserByID(dbStruct DBStructure, userID int) (*User, error) {
	for _, user := range dbStruct.Users {
		if user.ID == userID {
			return &user, nil
		}
	}
	return &User{}, errors.New("error finding user")
}

func SearchUserByEmail(dbStuct DBStructure, email string) *User {
	for _, user := range dbStuct.Users {
		if user.Email == email {
			return &user
		}
	}
	return nil
}

func createUserReponse(user User, token string) UserResponse {
	return UserResponse{
		ID: 	user.ID,
		Email: 	user.Email,
		Token: 	token,
	}
}
