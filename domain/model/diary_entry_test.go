package model

import (
	"github.com/suzan2go/familog-api/lib/token_generator"
	"mime/multipart"
	"testing"
)

func TestCreateDiaryEntry(t *testing.T) {

	db, cleanDB := InitTestDB(t)
	defer cleanDB("diary_entries")

	deviceToken := tokenGenerator.GenerateRandomToken(32)
	device, _ := db.FindOrCreateDeviceByToken(deviceToken)
	diary, _ := db.CreateDiary(&device.User, "日記帳")

	var initialCount int
	var afterCount int
	db.Table("diary_entries").Count(&initialCount)
	db.CreateDiaryEntry(&device.User, diary, "日記タイトル", "今日はこんなことしたよ", ":smile:", []*multipart.FileHeader{})
	db.Table("diary_entries").Count(&afterCount)

	if initialCount-afterCount != -1 {
		t.Error("diaries not created")
	}
}

func TestUpdateDiaryEntry(t *testing.T) {

	db, cleanDB := InitTestDB(t)
	defer cleanDB("diary_entries")

	deviceToken := tokenGenerator.GenerateRandomToken(32)
	device, _ := db.FindOrCreateDeviceByToken(deviceToken)
	diary, _ := db.CreateDiary(&device.User, "日記帳")
	diaryEntry, _ := db.CreateDiaryEntry(&device.User, diary, "日記タイトル", "今日はこんなことしたよ", ":smile:", []*multipart.FileHeader{})

	var initialCount int
	var afterCount int
	db.Table("diary_entries").Count(&initialCount)
	db.UpdateDiaryEntry(&device.User, diaryEntry, "更新後本文", diaryEntry.Body, diaryEntry.Emoji, []*multipart.FileHeader{})
	db.Table("diary_entries").Count(&afterCount)

	if initialCount-afterCount != 0 {
		t.Error("diary entry created")
	}

	if diaryEntry.Title != "更新後本文" {
		t.Error("diariy entry not updated")
	}
}
