package database

import (
	"os"
	"strconv"
)

// CreateChirp creates a new chirp and saves it to disk
func (db *DB) CreateChirp(body string, userID string) (Chirp, error) {
	dbStruct, err := db.loadDB()
	if err != nil {
		return Chirp{}, err
	}

	authorId, err := strconv.Atoi(userID)
	if err != nil {
		return Chirp{}, err
	}



	newID := len(dbStruct.Chirps) + 1

	newChirp := Chirp{
		ID: 	newID,
		Body: 	body,
		AuthorID: authorId,
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

func(db *DB) GetAllAuthorChirps(id int) ([]Chirp, error) {
	dbStruct, err := db.loadDB()
	if err != nil {
		return []Chirp{}, nil
	}

	chirpSlice := make([]Chirp, 0)
	for _, chirp := range dbStruct.Chirps {
		if chirp.AuthorID == id {
			chirpSlice = append(chirpSlice, chirp)
		}
	}
	return chirpSlice, nil
}