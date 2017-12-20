package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // for grom
	"os"
)

// DataStore implement database access func
type DataStore interface {
}

// DB struct
type DB struct {
	*gorm.DB
}

// InitDB intialize database connection
func InitDB() *DB {
	dbname := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")
	password := os.Getenv("DB_PASSWORD")
	user := os.Getenv("DB_USER")
	db, err := gorm.Open("postgres", "host="+host+" user="+user+" dbname="+dbname+" sslmode=disable password="+password)
	db.LogMode(true)
	if err != nil {
		panic("failed to connect database")
	}
	return &DB{db}
}

// Close close database connection
func (db *DB) Close() {
	db.Close()
}
