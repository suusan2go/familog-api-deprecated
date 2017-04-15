package model

import (
	"github.com/jinzhu/gorm"
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
	if err := tx.Create(diary).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	if err := tx.Create(&DiarySubscriber{UserID: user.ID, DiaryID: diary.ID}).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return diary, nil
}

// AllDiaries GetAllDiary
func (db *DB) AllDiaries(user *User) (*Diaries, error) {
	diaries := &Diaries{}
	if err := db.usersDiary(user).
		Find(&diaries.Diaries).Error; err != nil {
		return nil, err
	}
	return diaries, nil
}

// FindDiary find diary of id
func (db *DB) FindDiary(id string, user *User) (*Diary, error) {
	diary := &Diary{}
	if err := db.usersDiary(user).Find(&diary, "diaries.id = ?", id).Error; err != nil {
		return nil, err
	}
	return diary, nil
}

func (db *DB) usersDiary(user *User) *gorm.DB {
	return db.Joins("JOIN diary_subscribers on diary_subscribers.diary_id = diaries.id").
		Where("diary_subscribers.user_id = ?", user.ID)
}
