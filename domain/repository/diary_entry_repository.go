package repository

import (
	"github.com/suusan2go/familog-api/domain/model"
)

// DiaryEntryRepository repository interface for DiaryEntry
type DiaryEntryRepository interface {
	FindDiaryEntry(user *model.User, ID string) (*model.DiaryEntry, error)
}
