package gorm

import (
	"github.com/suusan2go/familog-api/domain/model"
)

// DiaryEntryImageRepository implemented by gorm
type DiaryEntryImageRepository struct {
	DB *model.DB
}

// Save persist DiaryEntry
func (repo DiaryEntryImageRepository) Save(diaryEntryImage *model.DiaryEntryImage) error {
	if diaryEntryImage.ID == 0 {
		return repo.DB.Create(diaryEntryImage).Error
	}
	return repo.DB.Update(diaryEntryImage).Error
}
