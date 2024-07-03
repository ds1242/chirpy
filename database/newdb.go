package database

import (
	// "os"
	"sync"
	// "encoding/json"
	// "fmt"
)

// NewDB creates a new database connection
// and creates the database file if it doesn't exist
func NewDB(path string) (*DB, error) {
	dbInstance := DB{
		path: path,
		mux: &sync.RWMutex{},
	}
	// Finish up ensureDB() here
	ensureErr := dbInstance.ensureDB()
	if ensureErr != nil {
		return &dbInstance, ensureErr
	}
	
	_, err := dbInstance.loadDB()
	if err != nil {
		return &dbInstance, err
	}
	

	return &dbInstance, nil
}