package repository

import (
	"github.com/suusan2go/familog-api/domain/model"
)

// DiaryEntryImageRepository repository interface for DiaryEntryImage
type DiaryEntryImageRepository interface {
	Save(diary *model.DiaryEntryImage) error
}
