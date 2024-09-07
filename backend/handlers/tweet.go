package handlers

import (
	"SupaTweet-backend/models"
	"encoding/json"
	"net/http"

	"gorm.io/gorm"
)

// PostTweet handles the creation of a new tweet.
func PostTweet(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var tweet models.Tweet

	// Decode the request body into the tweet model
	if err := json.NewDecoder(r.Body).Decode(&tweet); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Create the tweet in the database
	if err := db.Create(&tweet).Error; err != nil {
		http.Error(w, "Error creating tweet", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// GetTweets retrieves all tweets from the database.
func GetTweets(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var tweets []models.Tweet

	// Fetch all tweets from the database
	if err := db.Find(&tweets).Error; err != nil {
		http.Error(w, "Error retrieving tweets", http.StatusInternalServerError)
		return
	}

	// Encode the tweets into JSON and write the response
	if err := json.NewEncoder(w).Encode(tweets); err != nil {
		http.Error(w, "Error encoding tweets", http.StatusInternalServerError)
		return
	}
}
