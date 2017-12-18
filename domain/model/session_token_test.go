package model

import (
	"github.com/suzan2go/familog-api/lib/token_generator"
	"testing"
	"time"
)

func TestGenerateOrExtendSessionToken(t *testing.T) {
	db, cleanDB := InitTestDB(t)
	defer cleanDB("session_tokens")

	deviceToken := tokenGenerator.GenerateRandomToken(32)
	device, _ := db.FindOrCreateDeviceByToken(deviceToken)
	user := &device.User

	var (
		initialCount int
		afterCount   int
	)
	db.Table("session_tokens").Count(&initialCount)
	sessionToken, _ := db.GenerateOrExtendSessionToken(user)
	db.Table("session_tokens").Count(&afterCount)

	if afterCount-initialCount != 1 {
		t.Error("session token not generated")
	}

	currentSessionToken, _ := db.GenerateOrExtendSessionToken(user)

	if currentSessionToken.ExpiresAt.Unix() < sessionToken.ExpiresAt.Unix() {
		t.Error("session token expiresAt not extended")
	}
}

func TestFindSessionToken(t *testing.T) {
	db, cleanDB := InitTestDB(t)
	defer cleanDB("diary_entries")

	deviceToken := tokenGenerator.GenerateRandomToken(32)
	device, _ := db.FindOrCreateDeviceByToken(deviceToken)
	user := &device.User
	sessionToken, _ := db.GenerateOrExtendSessionToken(user)

	notExistToken, e1 := db.FindSessionToken("dummy")
	if notExistToken != nil && e1 != nil {
		t.Error("pass not existed token but return non nil value")
	}

	existToken, e2 := db.FindSessionToken(sessionToken.Token)
	if existToken.Token != sessionToken.Token {
		t.Error("Deifferent token returned")
	}
	if e2 != nil {
		t.Error("Some error returned", e2)
	}
}

func TestIsValid(t *testing.T) {
	db, cleanDB := InitTestDB(t)
	defer cleanDB("diary_entries")

	deviceToken := tokenGenerator.GenerateRandomToken(32)
	device, _ := db.FindOrCreateDeviceByToken(deviceToken)
	user := &device.User
	sessionToken, _ := db.GenerateOrExtendSessionToken(user)

	if sessionToken.IsValid() != true {
		t.Error("New session token returned but not valid", sessionToken.ExpiresAt)
	}

	sessionToken.ExpiresAt = time.Now().AddDate(0, -1, 0)
	if sessionToken.IsValid() != false {
		t.Error("Expired token returns true", sessionToken.ExpiresAt)
	}
}
