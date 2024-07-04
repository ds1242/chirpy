package database


func (db *DB) CreateUser(email string) (User, error) {
	dbStruct, err := db.loadDB()
	if err != nil {
		return User{}, err
	}

	newID := len(dbStruct.Users) + 1
	newUser := User{
		ID: 	newID,
		Email: 	email,
	}

	dbStruct.Users[newID] = newUser
	writeErr := db.writeDB(dbStruct)
	if writeErr != nil{
		return User{}, writeErr
	}
	return newUser, nil
}