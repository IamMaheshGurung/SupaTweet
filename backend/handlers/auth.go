package handlers

import (
	"SupaTweet-backend/models"
	"encoding/json"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("Its_Secret")

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func Register(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	myPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error in creating user", http.StatusInternalServerError)
		return
	}
	user.Password = string(myPassword)

	db.Create(&user)
	w.WriteHeader(http.StatusCreated)
}

func Login(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	var storedUser models.User
	db.Where("username = ?", user.Username).First(&storedUser)

	err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password),
		[]byte(user.Password))
	if err != nil {
		http.Error(w, "Invalid Credentials", http.StatusUnauthorized)
		return
	}
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
		http.Error(w, "error generating token", http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
}
