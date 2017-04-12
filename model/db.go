package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // for grom
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
	db, err := gorm.Open("postgres", "host=localhost user=postgres dbname=familog_development sslmode=disable password=password")
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
