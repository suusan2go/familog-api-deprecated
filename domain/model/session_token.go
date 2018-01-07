package model

import (
	"time"
)

// SessionToken model
type SessionToken struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	UserID    uint      `gorm:"not null;index" json:"userId"`
	Token     string    `gorm:"not null;unique_index" json:"token"`
	ExpiresAt time.Time `gorm:"not null" json:"expiresAt"`
	CreatedAt time.Time `gorm:"not null" json:"createdAt"`
	UpdatedAt time.Time `gorm:"not null" json:"updatedAt"`
	User      User      `json:"-"`
}

// IsValid valid token
func (s *SessionToken) IsValid() bool {
	return s.ExpiresAt.After(time.Now())
}
