package registry

import (
	"github.com/suusan2go/familog-api/domain/model"
	"github.com/suusan2go/familog-api/domain/repository"
	"github.com/suusan2go/familog-api/infrastructure/gorm"
)

// Registry DI
type Registry struct {
	DB *model.DB
}

// DiaryEntryRepository return DI
func (r Registry) DiaryEntryRepository() repository.DiaryEntryRepository {
	return gorm.DiaryEntryRepository{DB: r.DB}
}

// DiaryRepository return DI
func (r Registry) DiaryRepository() repository.DiaryRepository {
	return gorm.DiaryRepository{DB: r.DB}
}

// DeviceRepository return DI
func (r Registry) DeviceRepository() repository.DeviceRepository {
	return gorm.DeviceRepository{DB: r.DB}
}

// SessionRepository return DI
func (r Registry) SessionRepository() repository.SessionRepository {
	return gorm.SessionRepository{DB: r.DB}
}
