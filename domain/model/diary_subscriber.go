package model

import (
	"time"
)

// DiarySubscriber Model
type DiarySubscriber struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	UserID    uint      `gorm:"not null;index" json:"userID"`
	DiaryID   uint      `gorm:"not null;index" json:"diaryId"`
	CreatedAt time.Time `gorm:"not null" json:"createdAt"`
	UpdatedAt time.Time `gorm:"not null" json:"updatedAt"`
}
