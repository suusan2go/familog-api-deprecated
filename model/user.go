package model

import (
	"path/filepath"
	"time"

	"github.com/suzan2go/familog-api/lib/uploader"
)

// Image struct for image
type Image struct {
	URI  string `json:"uri"`
	Name string `json:"name"`
	Type string `json:"type"`
}

// User User model
type User struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Name      string    `json:"name"`
	Devices   []Device  `json:"-"`
	Image     Image     `gorm:"-" json:"image"`
	ImagePath string    `json:"-"`
	ImageURL  string    `gorm:"-" json:"-"`
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

// AfterFind gorm AfterFind callback implementation
func (u *User) AfterFind() (err error) {
	if len(u.ImagePath) == 0 {
		return
	}
	upl := uploader.InitUploader()
	url, err := upl.GetImageURL(u.ImagePath)
	if err != nil {
		return err
	}
	u.Image = Image{
		URI:  url.String(),
		Name: u.ImagePath,
		Type: "image/" + filepath.Ext(url.String()),
	}
	return
}
