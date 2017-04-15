package model

import (
	"time"
)

// DiaryEntry Model
type DiaryEntry struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	UserID    uint      `gorm:"not null;index" json:"userID"`
	DiaryID   uint      `gorm:"not null;index" json:"diaryId"`
	Title     string    `gorm:"not null" json:"title"`
	Body      string    `gorm:"not null" json:"body"`
	Emoji     string    `gorm:"not null" json:"emoji"`
	CreatedAt time.Time `gorm:"not null" json:"createdAt"`
	UpdatedAt time.Time `gorm:"not null" json:"updatedAt"`
	User      User
}

// CreateDiaryEntry create user related diary
func (db *DB) CreateDiaryEntry(
	user *User, diary *Diary, title string, body string, emoji string,
) (*DiaryEntry, error) {
	diaryEntry := &DiaryEntry{Title: title, Body: body, Emoji: emoji, DiaryID: diary.ID, UserID: user.ID}
	if err := db.Create(diaryEntry).Error; err != nil {
		return nil, err
	}
	return diaryEntry, nil
}
