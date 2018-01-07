package repository

import "github.com/suusan2go/familog-api/domain/model"

// SessionRepository repository interface for Session
type SessionRepository interface {
	GenerateOrExtendSessionToken(user *model.User) (*model.SessionToken, error)
	FindSessionToken(token string) (*model.SessionToken, error)
}
