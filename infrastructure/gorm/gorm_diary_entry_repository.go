package gorm

import (
	"github.com/jinzhu/gorm"

	"github.com/suusan2go/familog-api/domain/model"
)

// DiaryEntryRepository implemented by gorm
type DiaryEntryRepository struct {
	DB *model.DB
}

// AllDiaryEntries GetAllDiaryEntries
func (repo DiaryEntryRepository) AllDiaryEntries(diary *model.Diary) (*model.DiaryEntries, error) {
	diaryEntries := &model.DiaryEntries{}
	if err := repo.diaryScope(diary).
		Limit(10).
		Find(&diaryEntries.DiaryEntries).Error; err != nil {
		return nil, err
	}
	return diaryEntries, nil
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

// FindMyDiaryEntry find user wrote entry by id
func (repo DiaryEntryRepository) FindMyDiaryEntry(id string, user *model.User) (*model.DiaryEntry, error) {
	diaryEntry := &model.DiaryEntry{}
	if err := repo.myDiaryEntryScope(user).Preload("User").Preload("DiaryEntryImages").
		Find(diaryEntry, "diary_entries.id = ?", id).
		Error; err != nil {
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

func (repo DiaryEntryRepository) myDiaryEntryScope(user *model.User) *gorm.DB {
	return repo.DB.Where("diary_entries.user_id = ?", user.ID)
}

func (repo DiaryEntryRepository) diaryScope(diary *model.Diary) *gorm.DB {
	return repo.DB.Where("diary_entries.diary_id = ?", diary.ID).
		Order("diary_entries.id DESC").
		Preload("User").
		Preload("DiaryEntryImages")
}
