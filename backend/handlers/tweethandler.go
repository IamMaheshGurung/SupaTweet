package handlers

import (
	"SupaTweet-backend/database"
	"encoding/json"
	"log"
	"net/http"
)

type TweetHandler struct {
	L *log.Logger
}

func NewTweet(l *log.Logger) *TweetHandler {
	return &TweetHandler{l}
}

func (b TweetHandler) GetTweets(w http.ResponseWriter, r *http.Request) {
	t := database.GetTweets()
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(t)
	if err != nil {
		http.Error(w, "Unable to encode the tweet file", http.StatusInternalServerError)
	}
}

func (b TweetHandler) PostTweet(w http.ResponseWriter, r *http.Request) {
	var tweet database.Tweet
	json.NewDecoder(r.Body).Decode(&tweet)
	tweet.ID = len(database.GetTweets()) + 1
	database.PostTweets(&tweet)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(tweet)

}
