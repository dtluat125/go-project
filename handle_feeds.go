package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dtluat125/go-project/internal/database"
	"github.com/google/uuid"
)

func (apiConfig *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Title string `json:"title"`
		URL   string `json:"url"`
	}

	params := parameters{}
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request")
		return
	}

	feed, err := apiConfig.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:     uuid.New(),
		UserID: user.ID,
		Title:  params.Title,
		Url:    params.URL,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to create feed")
		return
	}

	respondWithJSON(w, http.StatusCreated, databaseFeedToAPIFeed(feed))
}

func (apiConfig *apiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := apiConfig.DB.GetFeeds(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to get feeds: %v", err))
		return
	}

	apiFeeds := make([]Feed, 0, len(feeds))
	for _, feed := range feeds {
		apiFeeds = append(apiFeeds, databaseFeedToAPIFeed(feed))
	}

	respondWithJSON(w, http.StatusOK, apiFeeds)
}
