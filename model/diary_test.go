package model

import (
	"github.com/suzan2go/familog-api/lib/token_generator"
	"testing"
)

func TestCreateDiary(t *testing.T) {

	db, cleanDB := InitTestDB(t)
	defer cleanDB("diaries")

	deviceToken := tokenGenerator.GenerateRandomToken(32)
	device, _ := db.FindOrCreateDeviceByToken(deviceToken)

	var initialCount int
	var afterCount int
	db.Table("diaries").Count(&initialCount)
	db.CreateDiary(&device.User, "ほげほげ")
	db.Table("diaries").Count(&afterCount)

	if initialCount-afterCount == 1 {
		t.Error("diaries created")
	}

	db.FindOrCreateDeviceByToken(deviceToken)

	if initialCount-afterCount == 0 {
		t.Error("diaries not created")
	}
}
