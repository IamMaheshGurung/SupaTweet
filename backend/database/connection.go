package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Connect establishes a connection to the database and returns a gorm.DB instance.
func Connect() (*gorm.DB, error) {
	var err error
	dsn := "host=localhost user=postgres password=newpassword dbname=supatweet port=5432 sslmode=disable"

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
		return nil, err
	}
	return DB, nil
}
