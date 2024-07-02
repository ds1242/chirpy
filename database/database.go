package database

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

type Chirp struct {
	ID 	 int 	`json: "id"`
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

// NewDB creates a new database connection
// and creates the database file if it doesn't exist
func NewDB(path string) (*DB, error) {
	dbInstance := DB{
		path: path,
		mux: &sync.RWMutex{},
	}
	data, err := os.ReadFile(dbInstance.path)
	if os.IsNotExist(err) {
		initialDB := DBStructure{Chirps: make(map[int]Chirp)}

		initialData, marshallErr := json.Marshal(initialDB)
		if marshallErr != nil {
			return nil, marshallErr
		}

		writeErr := os.WriteFile(dbInstance.path, initialData, 0644)
		if writeErr != nil {
			return nil, writeErr
		}
	} else if err != nil {
		return nil, err
	} else {
		var dbStruct DBStructure
		err = json.Unmarshal(data, &dbStruct)
		if err != nil {
			return nil, err
		}
		fmt.Printf("Loaded database: %v\n", dbStruct)
	}

	return &dbInstance, nil
}


// CreateChirp creates a new chirp and saves it to disk
func (db *DB) CreateChirp(body string) (Chirp, error) {
	db.mux.Lock()
	defer db.mux.Unlock()
	
	data, err := os.ReadFile(db.path)
    if err != nil {
        return Chirp{}, err
    }
	
	var dbStruct DBStructure
    if err := json.Unmarshal(data, &dbStruct); err != nil {
        return Chirp{}, err
    }

	dbStruct.LastID++
	newID := dbStruct.LastID

	newChirp := Chirp{
		ID: newID,
		Body: body,
	}
	
	dbStruct.Chirps[newID] = newChirp

	updatedData, marshallErr := json.Marshal(dbStruct)
	if marshallErr != nil {
		return Chirp{}, marshallErr
	}

	writeErr := os.WriteFile(db.path, updatedData, 0644)
	if writeErr != nil {
		return Chirp{}, writeErr
	}

	return newChirp, nil
}


// GetChirps returns all chirps in the database
// func (db *DB) GetChirps() ([]Chirp, error)


// ensureDB creates a new database file if it doesn't exist
// func (db *DB) ensureDB() error



// loadDB reads the database file into memory
// func (db *DB) loadDB() (DBStructure, error)


// writeDB writes the database file to disk
// func (db *DB) writeDB(dbStructure DBStructure) error