package model

import (
	"github.com/suzan2go/familog-api/uploader"
	"github.com/suzan2go/familog-api/util"
	"log"
	"mime/multipart"
	"path/filepath"
	"strconv"
	"time"
)

// DiaryEntryImage Model
type DiaryEntryImage struct {
	ID           uint      `gorm:"primary_key" json:"id"`
	DiaryEntryID uint      `gorm:"not null;index" json:"diaryEntryId"`
	FilePath     string    `gorm:"not null" json:"-"`
	FileURL      string    `json:"fileUrl"`
	CreatedAt    time.Time `gorm:"not null" json:"createdAt"`
	UpdatedAt    time.Time `gorm:"not null" json:"updatedAt"`
}

// CreateDiaryEntryImage create diary entry images
func (db *DB) CreateDiaryEntryImage(file *multipart.FileHeader, diaryEntry *DiaryEntry) (*DiaryEntryImage, error) {
	filePath := filepath.Join("diary_entry_images",
		strconv.Itoa(int(diaryEntry.ID)),
		util.GenerateRandomToken(16)+filepath.Ext(file.Filename),
	)
	diaryEntryImage := &DiaryEntryImage{DiaryEntryID: diaryEntry.ID, FilePath: filePath}
	db.Create(diaryEntryImage)
	if err := diaryEntryImage.UploadFile(file); err != nil {
		return nil, err
	}
	return diaryEntryImage, nil
}

// DeleteDiaryEntryImage delete uploaded file and database
func (db *DB) DeleteDiaryEntryImage(diaryEntryImage *DiaryEntryImage) error {
	if err := diaryEntryImage.DeleteFile(); err != nil {
		return err
	}
	if err := db.Delete(diaryEntryImage).Error; err != nil {
		return err
	}
	return nil
}

// UploadFile upload file
func (image *DiaryEntryImage) UploadFile(file *multipart.FileHeader) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	uploader := uploader.InitUploader()
	_, errr := uploader.UploadImage(src, image.FilePath)
	if errr != nil {
		return errr
	}
	return nil
}

// DeleteFile uploaded file
func (image *DiaryEntryImage) DeleteFile() error {
	uploader := uploader.InitUploader()
	err := uploader.DeleteImage(image.FilePath)
	if err != nil {
		return err
	}
	return nil
}

// AfterFind gorm AfterFind callback implementation
func (image *DiaryEntryImage) AfterFind() (err error) {
	u := uploader.InitUploader()
	url, err := u.GetImageURL(image.FilePath)
	if err != nil {
		return err
	}
	image.FileURL = url.String()
	log.Printf(image.FileURL)
	return
}
