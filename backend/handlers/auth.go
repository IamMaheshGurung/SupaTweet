package handlers

import (
	"SupaTweet-backend/models"
	"encoding/json"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var jwtKey = []byte("Its_Secret")

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func Register(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var user models.User

	// Decode request body
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Hash the password
	myPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error in hashing password", http.StatusInternalServerError)
		return
	}
	user.Password = string(myPassword)

	// Create the user in the database
	if err := db.Create(&user).Error; err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func Login(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var user models.User

	// Decode request body
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	var storedUser models.User

	// Find user in the database
	if err := db.Where("username = ?", user.Username).First(&storedUser).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, "Invalid Credentials", http.StatusUnauthorized)
			return
		}
		http.Error(w, "Error finding user", http.StatusInternalServerError)
		return
	}

	// Check password
	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password)); err != nil {
		http.Error(w, "Invalid Credentials", http.StatusUnauthorized)
		return
	}

	// Create JWT token
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	// Set cookie
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})

	w.WriteHeader(http.StatusOK)
}
