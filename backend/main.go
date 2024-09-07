package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"SupaTweet-backend/database"
	"SupaTweet-backend/handlers"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Tweet struct {
	ID      uint   `gorm:"primary_key"`
	Content string `gorm:"type:text"`
	UserID  uint
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&Tweet{})
}

func main() {
	// Initialize the router
	r := mux.NewRouter()

	// Connect to the database
	db, err := database.Connect()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	r.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		handlers.Register(db, w, r)
	}).Methods("POST")

	r.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		handlers.Login(db, w, r)
	}).Methods("POST")

	r.HandleFunc("/tweet", func(w http.ResponseWriter, r *http.Request) {
		handlers.PostTweet(db, w, r)
	}).Methods("POST")

	r.HandleFunc("/tweets", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetTweets(db, w, r)
	}).Methods("GET")

	// Configure the HTTP server
	srv := &http.Server{
		Handler:      r,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Start the server in a goroutine
	go func() {
		log.Println("Starting server on :8080")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error starting server: %v", err)
		}
	}()

	// Graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	<-stop

	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Error during server shutdown: %v", err)
	}

	log.Println("Server gracefully stopped")
}
