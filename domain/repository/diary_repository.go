package repository

import (
	"github.com/suusan2go/familog-api/domain/model"
)

// DiaryEntryRepository repository interface for DiaryEntry
type DiaryEntryRepository interface {
	FindDiaryEntries(diary model.Diary) (*model.DiaryEntries, error)
	saveDiaryEntry(diaryEntry model.DiaryEntry) (*model.DiaryEntry, error)
}
