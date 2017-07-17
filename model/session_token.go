package model

import (
	"github.com/suzan2go/familog-api/lib/token_generator"
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
func (s SessionToken) IsValid() bool {
	return s.ExpiresAt.After(time.Now())
}

// GenerateOrExtendSessionToken generate session token
func (db *DB) GenerateOrExtendSessionToken(user *User) (*SessionToken, error) {
	sessionToken := &SessionToken{}
	if err := db.Joins("JOIN users ON session_Tokens.user_id = users.id").
		Where("Session_Tokens.expires_At > ? AND users.id = ?", time.Now(), user.ID).
		FirstOrInit(sessionToken).Error; err != nil {
		return nil, err
	}
	if db.NewRecord(sessionToken) == false {
		return sessionToken, nil
	}
	sessionToken.Token = tokenGenerator.GenerateRandomToken(32)
	sessionToken.UserID = user.ID
	sessionToken.ExpiresAt = time.Now().AddDate(0, 1, 0)
	if err := db.Create(&sessionToken).Error; err != nil {
		return nil, err
	}
	return sessionToken, nil
}

// FindSessionToken find session token
func (db *DB) FindSessionToken(token string) (*SessionToken, error) {
	sessionToken := &SessionToken{}
	if err := db.Where("session_tokens.token = ?", token).Preload("User").First(&sessionToken).Error; err != nil {
		return nil, err
	}
	return sessionToken, nil
}
