package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	// "github.com/AanishRahmani/rssAggregator/internal/auth"

	"github.com/AanishRahmani/rssAggregator/internal/databases"

	// "github.com/AanishRahmani/rssAggregator/internal/databases"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user databases.User) {

	type parameters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("error parsing json: %v", err))
		return
	}

	feed, err := apiCfg.DB.CreateFeed(r.Context(), databases.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.Url,
		UserID:    user.ID,
	})
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("couldnt create feed: %v", err))
		return
	}

	responseWithJSON(w, 201, databaseFeedToFeed(feed))

}

func (apiCfg *apiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request) {

	feeds, err := apiCfg.DB.GetFeeds(r.Context())
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("couldnt get feed: %v", err))
		return
	}

	responseWithJSON(w, 201, databaseFeedsToFeeds(feeds))

}
