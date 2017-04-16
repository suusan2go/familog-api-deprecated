package model

import (
	"github.com/jinzhu/gorm"
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
	diaryEntry.User = *user
	return diaryEntry, nil
}

// UpdateDiaryEntry update user related diary
func (db *DB) UpdateDiaryEntry(
	user *User, id string, title string, body string, emoji string,
) (*DiaryEntry, error) {
	diaryEntry := &DiaryEntry{}
	if err := db.myDiaryScope(user).Find(diaryEntry, "diary_entries.id = ?", id).Error; err != nil {
		return nil, err
	}
	diaryEntry.Title = title
	diaryEntry.Body = body
	diaryEntry.Emoji = emoji
	if err := db.Save(diaryEntry).Error; err != nil {
		return nil, err
	}
	return diaryEntry, nil
}

// FindMyDiaryEntry find my diary entry
func (db *DB) FindMyDiaryEntry(id string, user *User) (*DiaryEntry, error) {
	diaryEntry := &DiaryEntry{}
	if err := db.myDiaryEntryScope(user).Preload("User").
		Find(diaryEntry, "diary_entries.id = ?", id).
		Error; err != nil {
		return nil, err
	}
	return diaryEntry, nil
}

func (db *DB) myDiaryEntryScope(user *User) *gorm.DB {
	return db.Where("diary_entries.user_id = ?", user.ID)
}
