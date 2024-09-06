package database

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var DB *gorm.DB

func Connect() {
	var err error
	DB, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=supatweet sslmode=disable password=yourpassword")
	if err != nil {
		log.Fatal("Failed to connect the database", err)

	}
	log.Println("DataBase is connected now")
}
