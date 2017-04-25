package model

import (
	"github.com/jinzhu/gorm"
	"os"
	"testing"
)

// InitTestDB setup DB for TESTING
func InitTestDB(t *testing.T) (DB, func(string)) {
	dbname := os.Getenv("TEST_DB_NAME")
	host := os.Getenv("DB_HOST")
	password := os.Getenv("DB_PASSWORD")
	user := os.Getenv("DB_USER")
	db, err := gorm.Open("postgres", "host="+host+" user="+user+" dbname="+dbname+" sslmode=disable password="+password)
	db.LogMode(true)
	if err != nil {
		t.Fatal("failed to connect database")
	}
	dbStruct := DB{db}
	dbStruct.Migration()
	return dbStruct, func(targetTable string) {
		db.DropTable(targetTable)
		db.Close()
	}
}
