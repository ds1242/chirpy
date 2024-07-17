package main

import (
	// "fmt"
	"net/http"
	"sort"
	"strconv"
	"fmt"

	"github.com/ds1242/chirpy/database"
	"github.com/ds1242/chirpy/helpers"
)

// This is the handler middleware func that queries all chirps by default but does accept an optional query parameter and returns all chirps for an author_id
func (cfg *apiConfig) GetAllChirpsHandler(w http.ResponseWriter, r *http.Request) {
	// initialize chirps and error variables
	dbChirps := make([]database.Chirp, 0)
	var dbChirpErr error

	// check for author_id optional parameter
	author_id := r.URL.Query().Get("author_id")
	
	// check for sort parameter
	sortParam := r.URL.Query().Get("sort")
	if len(sortParam) == 0 {
		sortParam = "asc"
	}
	if sortParam != "asc" || sortParam != "desc" {
		sortParam = "asc"
	}
	fmt.Println(sortParam)
	if len(author_id) == 0 {
		dbChirps, dbChirpErr = cfg.DB.GetChirps()
	} else {
		auth_id, err := strconv.Atoi(author_id)
		if err != nil {
			helpers.RespondWithError(w, http.StatusInternalServerError, "unable to parse author id")
			return
		}
		dbChirps, dbChirpErr = cfg.DB.GetAllAuthorChirps(auth_id)
	}

	if dbChirpErr != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "unable to fetch chirps")
		return
	}

	chirps := []database.Chirp{}
	for _, dbChirp := range dbChirps {
		chirps = append(chirps, database.Chirp{
			ID:       dbChirp.ID,
			Body:     dbChirp.Body,
			AuthorID: dbChirp.AuthorID,
		})
	}

	sort.Slice(chirps, func(i, j int) bool {
		return chirps[i].ID < chirps[j].ID
	})

	helpers.RespondWithJSON(w, http.StatusOK, chirps)
}

func (cfg *apiConfig) GetSingleChirpHandler(w http.ResponseWriter, r *http.Request) {
	chirpID, err := strconv.Atoi(r.PathValue("chirpID"))
	if err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, "Invalid chirp ID")
		return
	}
	chirp, err := cfg.DB.GetSingleChirp(chirpID)
	if err != nil {
		helpers.RespondWithError(w, http.StatusNotFound, "error finding chirp")
		return
	}
	helpers.RespondWithJSON(w, http.StatusOK, database.Chirp{
		ID:       chirp.ID,
		Body:     chirp.Body,
		AuthorID: chirp.AuthorID,
	})
}
