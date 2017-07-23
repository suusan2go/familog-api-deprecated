package model

import (
	"errors"
	"time"
)

// DiaryEntryComment diary entry related comments
type DiaryEntryComment struct {
	ID           uint      `gorm:"primary_key" json:"id"`
	UserID       uint      `gorm:"not null;index" json:"userID"`
	DiaryEntryID uint      `gorm:"not null;index" json:"diaryId"`
	Body         string    `gorm:"not null" json:"body"`
	CreatedAt    time.Time `gorm:"not null" json:"createdAt"`
	UpdatedAt    time.Time `gorm:"not null" json:"updatedAt"`
	User         User      `json:"user"`
}

// DiaryEntryComments diary entry comments
type DiaryEntryComments struct {
	DiaryEntries []DiaryEntryComment `json:"diaryEntryComments"`
}

// CreateDiaryEntryComment create user commented Comment
func (db *DB) CreateDiaryEntryComment(
	user *User, diaryEntry *DiaryEntry, body string,
) (*DiaryEntryComment, error) {
	diaryEntryComment := &DiaryEntryComment{Body: body, UserID: user.ID}
	if err := db.Create(diaryEntryComment).Error; err != nil {
		return nil, err
	}
	return diaryEntryComment, nil
}

// DeleteDiaryEntryComment delete user related Comment
func (db *DB) DeleteDiaryEntryComment(
	user *User, diaryEntryComment *DiaryEntryComment,
) error {
	if diaryEntryComment.UserID != user.ID {
		return errors.New("not user's comment")
	}
	if err := db.Delete(diaryEntryComment).Error; err != nil {
		return err
	}
	return nil
}
