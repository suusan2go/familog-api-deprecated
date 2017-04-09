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

// Diaries slice
type Diaries struct {
	Diaries []Diary `json:"diaries"`
}

// AllDiaries GetAllDiary
func (db *DB) AllDiaries() (*Diaries, error) {
	diaries := &Diaries{}
	if err := db.Find(&diaries.Diaries).Error; err != nil {
		return nil, err
	}
	return diaries, nil
}

// FindDiary find diary of id
func (db *DB) FindDiary(id string) (*Diary, error) {
	diary := &Diary{}
	if err := db.First(&diary, 1).Error; err != nil {
		return nil, err
	}
	return diary, nil
}
