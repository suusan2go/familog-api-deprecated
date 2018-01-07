package gorm

import (
	"testing"

	"github.com/suusan2go/familog-api/domain/model"
	"github.com/suusan2go/familog-api/lib/token_generator"
)

func TestGenerateOrExtendSessionToken(t *testing.T) {
	db, cleanDB := model.InitTestDB(t)
	repo := SessionRepository{DB: &db}
	defer cleanDB("session_tokens")

	deviceToken := tokenGenerator.GenerateRandomToken(32)
	device, _ := db.FindOrCreateDeviceByToken(deviceToken)
	user := &device.User

	var (
		initialCount int
		afterCount   int
	)
	db.Table("session_tokens").Count(&initialCount)
	sessionToken, _ := repo.GenerateOrExtendSessionToken(user)
	db.Table("session_tokens").Count(&afterCount)

	if afterCount-initialCount != 1 {
		t.Error("session token not generated")
	}

	currentSessionToken, _ := repo.GenerateOrExtendSessionToken(user)

	if currentSessionToken.ExpiresAt.Unix() < sessionToken.ExpiresAt.Unix() {
		t.Error("session token expiresAt not extended")
	}
}

func TestFindSessionToken(t *testing.T) {
	db, cleanDB := model.InitTestDB(t)
	repo := SessionRepository{DB: &db}
	defer cleanDB("diary_entries")

	deviceToken := tokenGenerator.GenerateRandomToken(32)
	device, _ := db.FindOrCreateDeviceByToken(deviceToken)
	user := &device.User
	sessionToken, _ := repo.GenerateOrExtendSessionToken(user)

	notExistToken, e1 := repo.FindSessionToken("dummy")
	if notExistToken != nil && e1 != nil {
		t.Error("pass not existed token but return non nil value")
	}

	existToken, e2 := repo.FindSessionToken(sessionToken.Token)
	if existToken.Token != sessionToken.Token {
		t.Error("Deifferent token returned")
	}
	if e2 != nil {
		t.Error("Some error returned", e2)
	}
}
