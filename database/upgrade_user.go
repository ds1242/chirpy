package database


func (db *DB) UpgradeUserToChirpyRed(userID int) error {
	// load dbStruct
	dbStruct, dbLoadErr := db.loadDB()
	
	if dbLoadErr != nil {
		return dbLoadErr
	}
	// search for user
	user, userErr := GetUserByID(dbStruct, userID)
	if userErr != nil {
		return userErr
	}
	// set user to chirpy red to true
	user.IsChirpyRed = true
	// update struct
	dbStruct.Users[userID] = *user
	// write to db
	writeErr := db.writeDB(dbStruct)
	if writeErr != nil{
		return writeErr
	}
	return nil
}