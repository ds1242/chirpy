package database

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"log"
	"strconv"
	"time"

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

	refreshToken, refreshDate, refreshTokenErr := GenerateRefreshToken() 
	if refreshTokenErr != nil {
		return UserResponse{}, refreshTokenErr
	}

	newID := len(dbStruct.Users) + 1
	newUser := User{
		ID: 			newID,
		Password: 		passHash,
		Email: 			email,
		RefreshToken: 	refreshToken,
		RefreshExpiration: refreshDate,
	}

	dbStruct.Users[newID] = newUser
	writeErr := db.writeDB(dbStruct)
	if writeErr != nil{
		return UserResponse{}, writeErr
	}

	expiresInSeconds := 60 * 60
	token, tokenErr := CreateToken(newUser.ID, expiresInSeconds, jwtSecret)

	if tokenErr != nil {
		return UserResponse{}, tokenErr
	}

	userResponse := createUserReponse(newUser, token)
	return userResponse, nil
}


func (db *DB) UserLogin(password string, email string, expiresInSeconds int, jwtSecret string) (UserResponse, error) {
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

	token, tokenErr := CreateToken(existingUser.ID, expiresInSeconds, jwtSecret)
	if tokenErr != nil {
		return UserResponse{}, tokenErr
	}
	// if existingUser.RefreshToken
	if time.Now().UTC().After(existingUser.RefreshExpiration) {		
		refreshToken, refreshDate, refreshTokenErr := GenerateRefreshToken() 
		if refreshTokenErr != nil {
			return UserResponse{}, refreshTokenErr
		}
		existingUser.RefreshToken = refreshToken
		existingUser.RefreshExpiration = refreshDate
		dbStruct.Users[existingUser.ID] = *existingUser

		writeErr := db.writeDB(dbStruct)
		if writeErr != nil{
			return UserResponse{}, writeErr
		}
	}
	
	userResponse := createUserReponse(*existingUser, token)
	return userResponse, nil
}



func (db *DB) UserUpdate(userID string, email string, password string, tokenString string) (UserResponse, error) {
	id, err := strconv.Atoi(userID)

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

	userResponse := createUserReponse(*user, tokenString)
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
		ID: 			user.ID,
		Email: 			user.Email,
		Token: 			token,
		RefreshToken: 	user.RefreshToken,
	}
}

func GenerateRefreshToken()(string, time.Time, error) {
	refreshToken := make([]byte, 32)
	_, err := rand.Read(refreshToken)
	if err != nil {
		return "", time.Time{}, err
	}

	encodedString := hex.EncodeToString(refreshToken)
	refreshDate := time.Now().Add(time.Duration(60) * 24 * time.Hour).UTC()

	return encodedString, refreshDate, nil
}
