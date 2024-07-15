package database

import (
	"os"
)

// CreateChirp creates a new chirp and saves it to disk
func (db *DB) CreateChirp(body string, jwtToken string) (Chirp, error) {
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