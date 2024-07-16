package database

import (
	"os"
	"errors"
)

func (db *DB) DeleteChirp(chirpID int, authorId int) error {
	dbStructure, err := db.loadDB()
	if err != nil {
		return err
	}

	chirp, ok := dbStructure.Chirps[chirpID]
	if !ok {
		return os.ErrNotExist
	}

	if chirp.AuthorID == authorId {
		delete(dbStructure.Chirps, chirpID)
		return nil
	}
	return errors.New("author id does not match")
}