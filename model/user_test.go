package model

import (
	"github.com/suzan2go/familog-api/util"
	"testing"
)

func TestFindUserByDeviceToken(t *testing.T) {
	db, cleanDB := InitTestDB(t)
	defer cleanDB("session_tokens")

	deviceToken := util.GenerateRandomToken(32)
	device, _ := db.FindOrCreateDeviceByToken(deviceToken)

	user1, e1 := db.FindUserByDeviceToken(device.Token)

	if user1 == nil {
		t.Error("exist token passed but user not found", e1)
	}

	user2, e2 := db.FindUserByDeviceToken("non exitst token")

	if user2 != nil {
		t.Error("non exist token passed but user found")
	}

	if user2 == nil && e2 == nil {
		t.Error("user not found but error is nil")
	}
}

func TestFindUserBySessionToken(t *testing.T) {
	db, cleanDB := InitTestDB(t)
	defer cleanDB("diary_entries")

	deviceToken := util.GenerateRandomToken(32)
	device, _ := db.FindOrCreateDeviceByToken(deviceToken)
	user := &device.User
	sessionToken, _ := db.GenerateOrExtendSessionToken(user)

	user1, e1 := db.FindUserBySessionToken(sessionToken.Token)

	if user1 == nil {
		t.Error("exist token passed but user not found", e1)
	}

	user2, e2 := db.FindUserBySessionToken("non exitst token")

	if user2 != nil {
		t.Error("non exist token passed but user found")
	}

	if user2 == nil && e2 == nil {
		t.Error("user not found but error is nil")
	}
}
