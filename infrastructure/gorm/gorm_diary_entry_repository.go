package gorm

import (
	"github.com/jinzhu/gorm"

	"github.com/suusan2go/familog-api/domain/model"
)

// DiaryEntryRepository implemented by gorm
type DiaryEntryRepository struct {
	DB *model.DB
}

// FindDiaryEntry find user subscribed diary entry by id
func (repo DiaryEntryRepository) FindDiaryEntry(user *model.User, ID string) (*model.DiaryEntry, error) {
	diaryEntry := &model.DiaryEntry{}
	if err := repo.subscribedDiaryEntryScope(user).
		Where("diary_entries.id = ?", ID).
		First(diaryEntry).Error; err != nil {
		return nil, err
	}
	return diaryEntry, nil
}

func (repo DiaryEntryRepository) subscribedDiaryEntryScope(user *model.User) *gorm.DB {
	return repo.DB.Joins("join diary_subscribers on diary_subscribers.diary_id = diary_entries.diary_id").
		Where("diary_subscribers.user_id = ?", user.ID).
		Preload("User").
		Preload("DiaryEntryImages")
}