package model

import (
	"time"
)

// Device Model
type Device struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Token     string    `gorm:"not null" json:"deviceToken"`
	UserID    int       `gorm:"not null;index" json:"userId"`
	CreatedAt time.Time `gorm:"not null" json:"createdAt"`
	UpdatedAt time.Time `gorm:"not null" json:"updatedAt"`
	User      User
}

// FindOrCreateDeviceByToken find or create device
func (db *DB) FindOrCreateDeviceByToken(deviceToken string) (*Device, error) {
	device := &Device{}
	if err := db.Where(&Device{Token: deviceToken}).FirstOrInit(&device).Error; err != nil {
		return nil, err
	}
	if device.UserID == 0 {
		user := &User{
			Devices: []Device{
				{Token: deviceToken},
			},
		}
		if err := db.Create(user).Error; err != nil {
			return nil, err
		}
	}
	return device, nil
}
