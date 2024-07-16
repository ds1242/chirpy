package database

import (
	"sync"
	"time"
)


type UserResponse struct {
	ID 				int 	`json:"id"`
	Email			string	`json:"email"`
	Token			string	`json:"token"`
	RefreshToken 	string	`json:"refresh_token"`
	IsChirpyRed		bool	`json:"is_chirpy_red"`
}
type Chirp struct {
	ID 			int 	`json:"id"`
	Body 		string 	`json:"body"`
	AuthorID 	int 	`json:"author_id"`
}

type User struct {
	ID 					int 		`json:"id"`
	Password 			[]byte		`json:"password"` 
	Email 				string 		`json:"email"`
	RefreshToken		string		`json:"refresh_token"`
	RefreshExpiration	time.Time	`json:"refresh_token_expiration"`
	IsChirpyRed			bool		`json:"is_chirpy_red"`
}

type UpdateUserParams struct {
    Email    string `json:"email,omitempty"`
    Password string `json:"password,omitempty"`
}


type DB struct {
	path string
	mux  *sync.RWMutex
}

type DBStructure struct {
	Chirps 	map[int]Chirp 	`json:"chirps"`
	Users 	map[int]User 	`json:"users"`
}
