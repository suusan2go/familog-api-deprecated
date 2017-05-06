package model

import (
	"time"
)

// User User model
type User struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Name      string    `json:"name"`
	Devices   []Device  `json:"-"`
	CreatedAt time.Time `gorm:"not null" json:"createdAt"`
	UpdatedAt time.Time `gorm:"not null" json:"updatedAt"`
}

// FindUserByDeviceToken find or create device
func (db *DB) FindUserByDeviceToken(deviceToken string) (*User, error) {
	user := &User{}
	if err := db.Joins("JOIN devices ON devices.user_id = users.id").Where("devices.Token = ?", deviceToken).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// FindUserBySessionToken find or create device
func (db *DB) FindUserBySessionToken(sessionToken string) (*User, error) {
	user := &User{}
	if err := db.Joins("JOIN session_tokens ON session_tokens.user_id = users.id").
		Where("session_tokens.Token = ? AND session_tokens.expires_at > ?", sessionToken, time.Now()).
		First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
