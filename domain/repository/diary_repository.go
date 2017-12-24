package repository

import "github.com/suusan2go/familog-api/domain/model"

// DiaryRepository repository interface for Diary
type DiaryRepository interface {
	FindDiary(id string, user *model.User) (*model.Diary, error)
	AllDiaries(user *model.User) (*model.Diaries, error)
}
