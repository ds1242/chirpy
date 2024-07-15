package database

// import (

// )


func (db *DB) RevokeRefreshToken(refreshToken string) (error) {
	// Load DB
	dbStruct, err := db.loadDB()
	if err != nil {
		return err
	}

	existingUser, existingUserErr := SearchByRefreshToken(dbStruct, refreshToken)
	if existingUserErr != nil {
		return existingUserErr
	}

	existingUser.RefreshToken = ""

	dbStruct.Users[existingUser.ID] = *existingUser
	writeErr := db.writeDB(dbStruct)
	if writeErr != nil{
		return writeErr
	}

	return nil
}