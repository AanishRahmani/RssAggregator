package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/AanishRahmani/rssAggregator/internal/database"
	"github.com/google/uuid"
)

// func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {

// 	type parameters struct {
// 		name string `json:"name"`
// 	}
// 	decoder := json.NewDecoder(r.Body)
// 	params := parameters{}
// 	err := decoder.Decode(&params)
// 	if err != nil {
// 		responseWithError(w, 400, fmt.Sprintf("error parsing json:%v", err))
// 		return
// 	}

// 	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
// 		ID:        uuid.New(),
// 		CreatedAt: time.Now().UTC(),
// 		UpdatedAt: time.Now().UTC(),
// 		Name:      params.name,
// 	})
// 	if err != nil {
// 		responseWithError(w, 400, fmt.Sprintf("couldnt create user: %v", err))
// 		return
// 	}

// 	responseWithJSON(w, 200, databaseUsertoUser(user))

// }
func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {

	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("error parsing json: %v", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name, // <-- Fix here
	})
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("couldnt create user: %v", err))
		return
	}

	responseWithJSON(w, 200, databaseUsertoUser(user))

}
