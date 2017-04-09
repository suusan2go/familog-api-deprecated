package model

import (
	"time"
)

// Diary Model
type Diary struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func AllDiary(db *DB) *Diary {
	diary := &Diary{}
	db.Find(diary)
	return diary
}

func FindDiary(db *DB, id string) *Diary {
	diary := &Diary{}
	db.First(diary, id)
	return diary
}
