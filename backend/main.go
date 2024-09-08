package main

import (
	"SupaTweet-backend/handlers"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {

	// Configure the HTTP server

	// Start the server in a goroutine
	var srv *http.Server

	go func() {
		l := log.New(os.Stdout, "tweet-api", log.LstdFlags)
		ts := handlers.NewTweet(l)

		// Initialize the router
		r := mux.NewRouter()

		r.HandleFunc("/tweet", ts.PostTweet).Methods(http.MethodPost)
		r.HandleFunc("/tweets", ts.GetTweets).Methods(http.MethodGet)

		c := cors.New(cors.Options{
			AllowedOrigins:   []string{"*"},
			AllowedMethods:   []string{http.MethodGet, http.MethodPost},
			AllowedHeaders:   []string{"Content-Type"}, // Corrected field name
			AllowCredentials: true,
		})

		srv = &http.Server{
			Handler:      c.Handler(r),
			Addr:         ":8080",
			WriteTimeout: 15 * time.Second,
			ReadTimeout:  15 * time.Second,
			IdleTimeout:  60 * time.Second,
		}

		log.Println("Starting server on :8080")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error starting server: %v", err)
		}
	}()

	// Graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	signal.Notify(stop, os.Kill)
	<-stop

	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Error during server shutdown: %v", err)
	}

	log.Println("Server gracefully stopped")
}
