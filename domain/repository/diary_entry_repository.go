package repository

import (
	"github.com/suusan2go/familog-api/domain/model"
)

// DiaryEntryRepository repository interface for DiaryEntry
type DiaryEntryRepository interface {
	Save(diary *model.DiaryEntry) error
	FindDiaryEntry(user *model.User, ID string) (*model.DiaryEntry, error)
	FindMyDiaryEntry(id string, user *model.User) (*model.DiaryEntry, error)
	AllDiaryEntries(diary *model.Diary) (*model.DiaryEntries, error)
}
