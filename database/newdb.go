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
	db := &DB{
		path: 	path,
		mux:  	&sync.RWMutex{},
	}
	err := db.ensureDB()
	return db, err
}