package repository

import "github.com/suusan2go/familog-api/domain/model"

// DeviceRepository repository interface for Device
type DeviceRepository interface {
	FindSubscribers(diaryEntry *model.DiaryEntry) ([]*model.Device, error)
}
