package model

import (
	"testing"
)

func TestCreateDiary(t *testing.T) {

	db, cleanDB := InitTestDB(t)
	defer cleanDB("diaries")

	device, _ := db.FindOrCreateDeviceByToken("hogehoge")

	var initialCount int
	var afterCount int
	db.Table("diaries").Count(&initialCount)
	db.CreateDiary(&device.User, "ほげほげ")
	db.Table("devices").Count(&afterCount)
	db.FindOrCreateDeviceByToken("hogehoge")
	db.Table("diaries").Count(&initialCount)

	if initialCount-afterCount == 1 {
		t.Error("diaries created")
	}

	db.FindOrCreateDeviceByToken("hogehoge")

	if initialCount-afterCount == 0 {
		t.Error("diaries not created")
	}
}
