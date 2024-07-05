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
	ID 			int 	`json:"id"`
	Password 	[]byte	`json:"-"` 
	Email 		string 	`json:"email"`
}
type DB struct {
	path string
	mux  *sync.RWMutex
}

type DBStructure struct {
	Chirps 	map[int]Chirp 	`json:"chirps"`
	Users 	map[int]User 	`json:"users"`
}


// NewDB creates a new database connection
// and creates the database file if it doesn't exist
func NewDB(path string) (*DB, error) {
	db := &DB{
		path: 	path,
		mux:  	&sync.RWMutex{},
	}
	err := db.ensureDB()
	return db, err
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