package database

import (
	"encoding/json"
	"fmt"
	// "net/http"
	"os"
	"sync"

	// "github.com/ds1242/chirpy/helpers"
)

type Chirp struct {
	ID 	 int 	`json:"id"`
	Body string `json:"body"`
}
type DB struct {
	path string
	mux  *sync.RWMutex
}

type DBStructure struct {
	LastID int 			 `json:"last_id"`
	Chirps map[int]Chirp `json:"chirps"`
}

// CreateChirp creates a new chirp and saves it to disk
func (db *DB) CreateChirp(body string) (Chirp, error) {
	db.mux.Lock()
	defer db.mux.Unlock()

	dbStruct, err := db.loadDB()
	if err != nil {
		return Chirp{}, err
	}

	dbStruct.LastID++
	newID := dbStruct.LastID

	newChirp := Chirp{
		ID: newID,
		Body: body,
	}
	
	dbStruct.Chirps[newID] = newChirp

	

	writeErr := db.writeDB(dbStruct)
	if writeErr != nil {
		return Chirp{}, writeErr
	}

	return newChirp, nil
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

// ensureDB creates a new database file if it doesn't exist
func (db *DB) ensureDB() error {
	
	_, err := os.ReadFile(db.path)
	if os.IsNotExist(err) {
		initialDB := DBStructure{Chirps: make(map[int]Chirp)}

		writeErr := db.writeDB(initialDB)

		if writeErr != nil {
			return writeErr
		}
	}
	return nil
}



// loadDB reads the database file into memory
func (db *DB) loadDB() (DBStructure, error) {
	db.mux.RLock()
	defer db.mux.RUnlock()

	var dbStruct DBStructure
	err := db.ensureDB()
	if err != nil {
		return DBStructure{}, err
	}

	data, readErr := os.ReadFile(db.path)
	if readErr != nil {
		return DBStructure{}, readErr
	}
		
	err = json.Unmarshal(data, &dbStruct)
	if err != nil {
		return DBStructure{}, err
	}
	fmt.Printf("Loaded database: %v\n", dbStruct)

	return dbStruct, nil
	
}


// writeDB writes the database file to disk
func (db *DB) writeDB(dbStructure DBStructure) error {

	data, marshallErr := json.Marshal(dbStructure)
	if marshallErr != nil {
		return  marshallErr
	}

	writeErr := os.WriteFile(db.path, data, 0644)
	if writeErr != nil {
		return writeErr
	}
	return nil
}