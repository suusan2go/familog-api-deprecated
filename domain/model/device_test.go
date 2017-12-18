package model

import (
	"github.com/suzan2go/familog-api/lib/token_generator"
	"testing"
)

func TestFindOrCreateDeviceByToken(t *testing.T) {

	db, cleanDB := InitTestDB(t)
	defer cleanDB("devices")

	deviceToken := tokenGenerator.GenerateRandomToken(32)
	var initialCount int
	var afterCount int
	db.Table("devices").Count(&initialCount)

	device, firstErr := db.FindOrCreateDeviceByToken(deviceToken)
	db.Table("devices").Count(&afterCount)
	if firstErr != nil {
		t.Error("error has occuered when Device created")
	}
	if device.Token != deviceToken {
		t.Error("WrongValue setted")
	}
	if initialCount-afterCount != -1 {
		t.Error("devices not created")
	}

	createdDevice, sErr := db.FindOrCreateDeviceByToken(deviceToken)
	db.Table("devices").Count(&initialCount)
	if sErr != nil {
		t.Error("error has occuered when Device created")
	}
	if device.ID != createdDevice.ID {
		t.Error("not matched to first device")
	}
	if initialCount-afterCount != 0 {
		t.Error("devices created")
	}
}
