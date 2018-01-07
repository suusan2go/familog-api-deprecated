package gorm

import (
	"github.com/suusan2go/familog-api/domain/model"

	"time"

	"github.com/suusan2go/familog-api/lib/token_generator"
)

// SessionRepository implemented by gorm
type SessionRepository struct {
	DB *model.DB
}

// GenerateOrExtendSessionToken generate session token
func (repo SessionRepository) GenerateOrExtendSessionToken(user *model.User) (*model.SessionToken, error) {
	sessionToken := &model.SessionToken{}
	if err := repo.DB.Joins("JOIN users ON session_Tokens.user_id = users.id").
		Where("Session_Tokens.expires_At > ? AND users.id = ?", time.Now(), user.ID).
		FirstOrInit(sessionToken).Error; err != nil {
		return nil, err
	}
	if repo.DB.NewRecord(sessionToken) == false {
		return sessionToken, nil
	}
	sessionToken.Token = tokenGenerator.GenerateRandomToken(32)
	sessionToken.UserID = user.ID
	sessionToken.ExpiresAt = time.Now().AddDate(0, 1, 0)
	if err := repo.DB.Create(&sessionToken).Error; err != nil {
		return nil, err
	}
	return sessionToken, nil
}

// FindSessionToken find session token
func (repo SessionRepository) FindSessionToken(token string) (*model.SessionToken, error) {
	sessionToken := &model.SessionToken{}
	if err := repo.DB.Where("session_tokens.token = ?", token).Preload("User").First(&sessionToken).Error; err != nil {
		return nil, err
	}
	return sessionToken, nil
}
