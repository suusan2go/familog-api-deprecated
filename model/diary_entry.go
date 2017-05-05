package model

import (
	"github.com/jinzhu/gorm"
	"mime/multipart"
	"strconv"
	"time"
)

// DiaryEntry Model
type DiaryEntry struct {
	ID               uint              `gorm:"primary_key" json:"id"`
	UserID           uint              `gorm:"not null;index" json:"userID"`
	DiaryID          uint              `gorm:"not null;index" json:"diaryId"`
	Title            string            `gorm:"not null" json:"title"`
	Body             string            `gorm:"not null" json:"body"`
	Emoji            string            `gorm:"not null" json:"emoji"`
	CreatedAt        time.Time         `gorm:"not null" json:"createdAt"`
	UpdatedAt        time.Time         `gorm:"not null" json:"updatedAt"`
	User             User              `json:"user"`
	DiaryEntryImages []DiaryEntryImage `json:"diaryEntryImages"`
}

// DiaryEntries diary entries slice
type DiaryEntries struct {
	DiaryEntries []DiaryEntry `json:"diary_entries"`
}

// CreateDiaryEntry create user related diary
func (db *DB) CreateDiaryEntry(
	user *User, diary *Diary, title string, body string, emoji string,
	images []*multipart.FileHeader,
) (*DiaryEntry, error) {
	diaryEntry := &DiaryEntry{Title: title, Body: body, Emoji: emoji, DiaryID: diary.ID, UserID: user.ID}
	if err := db.Create(diaryEntry).Error; err != nil {
		return nil, err
	}
	for _, image := range images {
		if image != nil {
			if _, err := db.CreateDiaryEntryImage(image, diaryEntry); err != nil {
				return nil, err
			}
		}
	}
	d, err := db.FindMyDiaryEntry(strconv.Itoa(int(diaryEntry.ID)), user)
	if err != nil {
		return nil, err
	}
	return d, nil
}

// UpdateDiaryEntry update user related diary
func (db *DB) UpdateDiaryEntry(
	user *User, diaryEntry *DiaryEntry, title string, body string, emoji string,
) error {
	diaryEntry.Title = title
	diaryEntry.Body = body
	diaryEntry.Emoji = emoji
	if err := db.Save(diaryEntry).Error; err != nil {
		return err
	}
	return nil
}

// FindMyDiaryEntry find my diary entry
func (db *DB) FindMyDiaryEntry(id string, user *User) (*DiaryEntry, error) {
	diaryEntry := &DiaryEntry{}
	if err := db.myDiaryEntryScope(user).Preload("User").Preload("DiaryEntryImages").
		Find(diaryEntry, "diary_entries.id = ?", id).
		Error; err != nil {
		return nil, err
	}
	return diaryEntry, nil
}

// FindMyDiaryEntryImage find my diary entry
func (db *DB) FindMyDiaryEntryImage(diaryEntryID string, diaryEntryImageID string, user *User) (
	*DiaryEntryImage, error) {
	diaryEntry, err := db.FindMyDiaryEntry(diaryEntryID, user)
	if err != nil {
		return nil, err
	}
	diaryEntryImage := &DiaryEntryImage{}
	if err := db.Where("diary_entry_images.diary_entry_id = ?", diaryEntry.ID).
		Find(diaryEntryImage, "diary_entry_images.id = ?", diaryEntryImageID).Error; err != nil {
		return nil, err
	}
	return diaryEntryImage, nil
}

// AllDiaryEntries GetAllDiaryEntries
func (db *DB) AllDiaryEntries(user *User) (*DiaryEntries, error) {
	diaryEntries := &DiaryEntries{}
	if err := db.subscribedDiaryEntryScope(user).
		Limit(10).
		Find(&diaryEntries.DiaryEntries).Error; err != nil {
		return nil, err
	}
	return diaryEntries, nil
}

// MoreOlderDiaryEntries GetAllDiaryEntries id < sinceID
func (db *DB) MoreOlderDiaryEntries(user *User, sinceID string) (*DiaryEntries, error) {
	diaryEntries := &DiaryEntries{}
	if err := db.subscribedDiaryEntryScope(user).
		Limit(10).
		Where("diary_entries.id < ?", sinceID).
		Find(&diaryEntries.DiaryEntries).Error; err != nil {
		return nil, err
	}
	return diaryEntries, nil
}

// MoreNewerDiaryEntries GetAllDiaryEntries id > maxID
func (db *DB) MoreNewerDiaryEntries(user *User, maxID string) (*DiaryEntries, error) {
	diaryEntries := &DiaryEntries{}
	if err := db.subscribedDiaryEntryScope(user).
		Limit(10).
		Where("diary_entries.id > ?", maxID).
		Find(&diaryEntries.DiaryEntries).Error; err != nil {
		return nil, err
	}
	return diaryEntries, nil
}

func (db *DB) subscribedDiaryEntryScope(user *User) *gorm.DB {
	return db.Joins("JOIN diaries on diary_entries.diary_id = diaries.id").
		Joins("JOIN diary_subscribers on diary_subscribers.diary_id = diaries.id").
		Preload("User").
		Preload("DiaryEntryImages").
		Order("diary_entries.id DESC").
		Where("diary_subscribers.user_id = ?", user.ID)
}

func (db *DB) myDiaryEntryScope(user *User) *gorm.DB {
	return db.Where("diary_entries.user_id = ?", user.ID)
}
