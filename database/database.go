package database

import (
	"encoding/json"
	"fmt"
	"errors"
	"os"
	"sync"

	// "github.com/ds1242/chirpy/helpers"
)

type Chirp struct {
	ID 	 int 	`json:"id"`
	Body string `json:"body"`
}

type User struct {
	ID 		int 	`json:id`
	Email 	string 	`json:"email"`
}
type DB struct {
	path string
	mux  *sync.RWMutex
}

type DBStructure struct {
	Chirps 	map[int]Chirp 	`json:"chirps"`
	Users 	map[int]User 	`json:"users"`
}

// CreateChirp creates a new chirp and saves it to disk
func (db *DB) CreateChirp(body string) (Chirp, error) {
	dbStruct, err := db.loadDB()
	if err != nil {
		return Chirp{}, err
	}

	newID := len(dbStruct.Chirps) + 1

	newChirp := Chirp{
		ID: 	newID,
		Body: 	body,
	}
	
	dbStruct.Chirps[newID] = newChirp

	

	writeErr := db.writeDB(dbStruct)
	if writeErr != nil {
		return Chirp{}, writeErr
	}

	return newChirp, nil
}

func (db *DB) CreateUser(email string) (User, error) {
	dbStruct, err := db.loadDB()
	if err != nil {
		return User{}, err
	}

	newID := len(dbStruct.Users) + 1
	newUser := User{
		ID: 	newID,
		Email: 	email,
	}

	dbStruct.Users[newID] = newUser
	writeErr := db.writeDB(dbStruct)
	if writeErr != nil{
		return User{}, writeErr
	}
	return newUser, nil
}
// GetChirps returns all chirps in the database
func (db *DB) GetChirps() ([]Chirp, error) {
	dbStruct, err := db.loadDB()
	if err != nil {
		return nil, err
	}
	chirpSlice := make([]Chirp, 0, len(dbStruct.Chirps))
	for _, chirp := range dbStruct.Chirps {
		chirpSlice = append(chirpSlice, chirp)
	}
	return chirpSlice, nil
}

// Get Single Chirp from the database
func(db *DB) GetSingleChirp(id int) (Chirp, error) {
	dbStructure, err := db.loadDB()
	if err != nil {
		return Chirp{}, err
	}

	chirp, ok := dbStructure.Chirps[id]
	if !ok {
		return Chirp{}, os.ErrNotExist
	}
	return chirp, nil
}

func (db *DB) createDB() error {
	dbStructure := DBStructure{
		Chirps: map[int]Chirp{},
		Users: 	map[int]User{},
	}
	return db.writeDB(dbStructure)
}

// ensureDB creates a new database file if it doesn't exist
func (db *DB) ensureDB() error {	
	_, err := os.ReadFile(db.path)
	if errors.Is(err, os.ErrNotExist) {
		return db.createDB()
	}
	return err
}

// loadDB reads the database file into memory
func (db *DB) loadDB() (DBStructure, error) {
	db.mux.RLock()
	defer db.mux.RUnlock()

	dbStruct := DBStructure{}
	dat, err := os.ReadFile(db.path)
	if errors.Is(err, os.ErrNotExist) {
		return dbStruct, err
	}
	err = json.Unmarshal(dat, &dbStruct)
	if err != nil {
		return dbStruct, err
	}
	fmt.Printf("Loaded database: %v\n", dbStruct)

	return dbStruct, nil
	
}


// writeDB writes the database file to disk
func (db *DB) writeDB(dbStructure DBStructure) error {
	db.mux.Lock()
	defer db.mux.Unlock()

	data, marshallErr := json.Marshal(dbStructure)
	if marshallErr != nil {
		return  marshallErr
	}

	writeErr := os.WriteFile(db.path, data, 0600)
	if writeErr != nil {
		return writeErr
	}
	return nil
}