package handlers

import (
	"SupaTweet-backend/models"
	"encoding/json"
	"net/http"

	"github.com/jinzhu/gorm"
)

func PostTweet(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var tweet models.Tweet
	json.NewDecoder(r.Body).Decode(&tweet)
	db.Create(&tweet)
	w.WriteHeader(http.StatusCreated)
}

func GetTweets(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var tweets []models.Tweet
	db.Find(&tweets)
	json.NewEncoder(w).Encode(tweets)
}
