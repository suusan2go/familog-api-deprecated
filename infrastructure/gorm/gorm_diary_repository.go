package gorm

import (
	"github.com/jinzhu/gorm"
	"github.com/suusan2go/familog-api/domain/model"
)

// DiaryRepository implemented by gorm
type DiaryRepository struct {
	DB *model.DB
}

// FindDiary find user subscribed diary entry by id
func (repo DiaryRepository) FindDiary(id string, user *model.User) (*model.Diary, error) {
	diary := &model.Diary{}
	if err := repo.myDiaryScope(user).Find(&diary, "diaries.id = ?", id).Error; err != nil {
		return nil, err
	}
	return diary, nil
}

// AllDiaries GetAllDiary
func (repo DiaryRepository) AllDiaries(user *model.User) (*model.Diaries, error) {
	diaries := &model.Diaries{}
	if err := repo.myDiaryScope(user).
		Find(&diaries.Diaries).Error; err != nil {
		return nil, err
	}
	return diaries, nil
}

func (repo DiaryRepository) myDiaryScope(user *model.User) *gorm.DB {
	return repo.DB.Joins("JOIN diary_subscribers on diary_subscribers.diary_id = diaries.id").
		Where("diary_subscribers.user_id = ?", user.ID)
}
