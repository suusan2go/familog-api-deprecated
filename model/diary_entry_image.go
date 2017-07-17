package model

import (
	"github.com/suzan2go/familog-api/lib/token_generator"
	"github.com/suzan2go/familog-api/lib/uploader"
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
	FileURL      string    `gorm:"-" json:"uri"`
	CreatedAt    time.Time `gorm:"not null" json:"createdAt"`
	UpdatedAt    time.Time `gorm:"not null" json:"updatedAt"`
}

// CreateDiaryEntryImage create diary entry images
func (db *DB) CreateDiaryEntryImage(file *multipart.FileHeader, diaryEntry *DiaryEntry) (*DiaryEntryImage, error) {
	filePath := filepath.Join("diary_entry_images",
		strconv.Itoa(int(diaryEntry.ID)),
		tokenGenerator.GenerateRandomToken(16)+filepath.Ext(file.Filename),
	)
	diaryEntryImage := &DiaryEntryImage{DiaryEntryID: diaryEntry.ID, FilePath: filePath}
	db.Create(diaryEntryImage)
	if err := diaryEntryImage.UploadFile(file); err != nil {
		return nil, err
	}
	return diaryEntryImage, nil
}

// UpdateDiaryEntryImage create diary entry images
func (db *DB) UpdateDiaryEntryImage(file *multipart.FileHeader, diaryEntryImage *DiaryEntryImage) error {
	originalDiaryEntryImage := *diaryEntryImage
	filePath := filepath.Join("diary_entry_images",
		strconv.Itoa(int(diaryEntryImage.DiaryEntryID)),
		tokenGenerator.GenerateRandomToken(16)+filepath.Ext(file.Filename),
	)
	diaryEntryImage.DeleteFile()
	diaryEntryImage.FilePath = filePath
	if err := db.Save(diaryEntryImage).Error; err != nil {
		return err
	}
	if err := diaryEntryImage.UploadFile(file); err != nil {
		return err
	}
	originalDiaryEntryImage.DeleteFile()
	return nil
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
	upl := uploader.NewUploader()
	_, errr := upl.UploadImage(src, image.FilePath)
	if errr != nil {
		return errr
	}
	return nil
}

// DeleteFile uploaded file
func (image *DiaryEntryImage) DeleteFile() error {
	upl := uploader.NewUploader()
	if err := upl.DeleteImage(image.FilePath); err != nil {
		return err
	}
	return nil
}

// AfterFind gorm AfterFind callback implementation
func (image *DiaryEntryImage) AfterFind() (err error) {
	u := uploader.NewUploader()
	url, err := u.GetImageURL(image.FilePath)
	if err != nil {
		return err
	}
	image.FileURL = url.String()
	return
}
