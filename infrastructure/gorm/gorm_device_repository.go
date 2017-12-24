package gorm

import (
	"github.com/suusan2go/familog-api/domain/model"
)

// DeviceRepository implemented by gorm
type DeviceRepository struct {
	DB *model.DB
}

// FindSubscribers find user subscribed diaries
func (repo DiaryRepository) FindSubscribers(diaryEntry *model.DiaryEntry) ([]*model.Device, error) {
	var devices []*model.Device
	if err := repo.DB.Where("diary_subscribers.diary_id =?", diaryEntry.DiaryID).
		Where("diary_subscribers.user_id != ?", diaryEntry.UserID).
		Joins("JOIN users ON users.id = devices.user_id").
		Joins("JOIN diary_subscribers ON diary_subscribers.user_id = users.id").
		Find(&devices).Error; err != nil {
		return nil, err
	}
	return devices, nil
}
