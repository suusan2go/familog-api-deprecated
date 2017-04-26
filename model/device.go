package model

import (
	"time"
)

// Device Model
type Device struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Token     string    `gorm:"not null" json:"deviceToken"`
	UserID    uint      `gorm:"not null;index" json:"userId"`
	CreatedAt time.Time `gorm:"not null" json:"createdAt"`
	UpdatedAt time.Time `gorm:"not null" json:"updatedAt"`
	User      User      `json:"-"`
}

// FindOrCreateDeviceByToken find or create device
func (db *DB) FindOrCreateDeviceByToken(deviceToken string) (*Device, error) {
	device := &Device{}
	if err := db.Where(&Device{Token: deviceToken}).FirstOrInit(&device).Error; err != nil {
		return nil, err
	}
	if device.UserID == 0 {
		user := &User{}
		if err := db.Create(user).Error; err != nil {
			return nil, err
		}
		device.User = *user
		device.UserID = user.ID
		device.Token = deviceToken
		db.Create(device)
	}
	return device, nil
}
