package model

import (
	"time"
)

// Diary Model
type Diary struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Title     string    `gorm:"not null" json:"title"`
	CreatedAt time.Time `gorm:"not null" json:"createdAt"`
	UpdatedAt time.Time `gorm:"not null" json:"updatedAt"`
}

// Diaries slice
type Diaries struct {
	Diaries []Diary `json:"diaries"`
}

// CreateDiary create user related diary
func (db *DB) CreateDiary(user *User, title string) (*Diary, error) {
	diary := &Diary{Title: title}
	tx := db.Begin()
	txDB := &DB{tx}
	if err := txDB.Create(diary).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	// add subscriber
	if err := txDB.Create(&DiarySubscriber{UserID: user.ID, DiaryID: diary.ID}).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	// add invitation
	if _, err := txDB.CreateDiaryInvitation(diary); err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return diary, nil
}
