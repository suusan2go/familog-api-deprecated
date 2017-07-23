package model

import (
	"github.com/suzan2go/familog-api/lib/token_generator"
	"mime/multipart"
	"testing"
)

func TestCreateDiaryEntryComment(t *testing.T) {

	db, cleanDB := InitTestDB(t)
	defer cleanDB("diary_entry_comments")
	defer cleanDB("diary_entries")
	defer cleanDB("diaries")
	defer cleanDB("devices")

	deviceToken := tokenGenerator.GenerateRandomToken(32)
	device, _ := db.FindOrCreateDeviceByToken(deviceToken)
	diary, _ := db.CreateDiary(&device.User, "日記帳")
	diaryEntry, _ := db.CreateDiaryEntry(&device.User, diary, "日記タイトル", "今日はこんなことしたよ", ":smile:", []*multipart.FileHeader{})

	var initialCount int
	var afterCount int
	db.Table("diary_entry_comments").Count(&initialCount)
	diaryEntryComment, err := db.CreateDiaryEntryComment(&device.User, diaryEntry, "コメントテスト")
	db.Table("diary_entry_comments").Count(&afterCount)

	if initialCount-afterCount != -1 {
		t.Error("diary_entry_comment not created")
	}
	if err != nil {
		t.Errorf("diary_entry_comment not created %s", err)
	}
	if diaryEntryComment.Body != "コメントテスト" {
		t.Errorf("diary_entry_comment not saved %v", diaryEntryComment)
	}
}

func TestDeleteDiaryEntryComment(t *testing.T) {

	db, cleanDB := InitTestDB(t)
	defer cleanDB("diary_entry_comments")
	defer cleanDB("diary_entries")
	defer cleanDB("diaries")
	defer cleanDB("devices")

	deviceToken := tokenGenerator.GenerateRandomToken(32)
	device, _ := db.FindOrCreateDeviceByToken(deviceToken)
	diary, _ := db.CreateDiary(&device.User, "日記帳")
	diaryEntry, _ := db.CreateDiaryEntry(&device.User, diary, "日記タイトル", "今日はこんなことしたよ", ":smile:", []*multipart.FileHeader{})
	diaryEntryComment, _ := db.CreateDiaryEntryComment(&device.User, diaryEntry, "コメントテスト")

	var initialCount int
	var afterCount int
	db.Table("diary_entry_comments").Count(&initialCount)
	err := db.DeleteDiaryEntryComment(&device.User, diaryEntryComment)
	db.Table("diary_entry_comments").Count(&afterCount)

	if initialCount-afterCount != 1 {
		t.Error("diary entry not deleted %v", err)
	}
}
