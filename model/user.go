package model

import (
	"mime/multipart"
	"path/filepath"
	"strconv"
	"time"

	"github.com/suzan2go/familog-api/lib/uploader"
	"github.com/suzan2go/familog-api/util"
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

// UpdateUserImage create user images
func (db *DB) UpdateUserImage(file *multipart.FileHeader, user *User) error {
	originalUser := *user
	filePath := filepath.Join("users",
		strconv.Itoa(int(user.ID)),
		util.GenerateRandomToken(16)+filepath.Ext(file.Filename),
	)
	user.DeleteFile()
	user.ImagePath = filePath
	if err := db.Save(user).Error; err != nil {
		return err
	}
	if err := user.UploadFile(file); err != nil {
		return err
	}
	originalUser.DeleteFile()
	return nil
}

// AfterFind gorm AfterFind callback implementation
func (user *User) AfterFind() (err error) {
	if len(user.ImagePath) == 0 {
		return
	}
	upl := uploader.InitUploader()
	url, err := upl.GetImageURL(user.ImagePath)
	if err != nil {
		return err
	}
	user.Image = Image{
		URI:  url.String(),
		Name: user.ImagePath,
		Type: "image/" + filepath.Ext(url.String()),
	}
	return
}

// UploadFile upload file
func (user *User) UploadFile(file *multipart.FileHeader) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	uploader := uploader.InitUploader()
	_, errr := uploader.UploadImage(src, user.ImagePath)
	if errr != nil {
		return errr
	}
	return nil
}

// DeleteFile uploaded file
func (user *User) DeleteFile() error {
	uploader := uploader.InitUploader()
	if err := uploader.DeleteImage(user.ImagePath); err != nil {
		return err
	}
	return nil
}
