package model

import (
	"time"

	"github.com/suzan2go/familog-api/util"
)

// DiaryInvitation Model
type DiaryInvitation struct {
	ID             uint       `gorm:"primary_key" json:"id"`
	DiaryID        uint       `gorm:"not null" json:"diaryId"`
	InvitationCode string     `gorm:"not null" json:"invitationCode"`
	ExpiredAt      *time.Time `json:"expiredAt"`
	CreatedAt      time.Time  `gorm:"not null" json:"createdAt"`
	UpdatedAt      time.Time  `gorm:"not null" json:"updatedAt"`
}

// CreateDiaryInvitation create user related diary invitation
func (db *DB) CreateDiaryInvitation(diary *Diary) (*DiaryInvitation, error) {
	diaryInvitation := &DiaryInvitation{DiaryID: diary.ID, InvitationCode: util.GenerateRandomToken(32), ExpiredAt: nil}
	if err := db.Create(diaryInvitation).Error; err != nil {
		return nil, err
	}
	return diaryInvitation, nil
}

// FindNotExpiredDiaryInvitation get current not expired invitation code
func (db *DB) FindNotExpiredDiaryInvitation(diary *Diary) (*DiaryInvitation, error) {
	diaryInvitation := &DiaryInvitation{}
	if err := db.Where("diary_id = ? AND expired_at is NULL", diary.ID).First(diaryInvitation).Error; err != nil {
		return nil, err
	}
	return diaryInvitation, nil
}

// RecreateDiaryInvitation expire current invitation and create user related diary invitation
func (db *DB) RecreateDiaryInvitation(diary *Diary) (*DiaryInvitation, error) {
	currentDiaryInvitation, _ := db.FindNotExpiredDiaryInvitation(diary)
	if currentDiaryInvitation != nil {
		db.Model(currentDiaryInvitation).Update("ExpiredAt", time.Now())
	}
	diaryInvitation, err := db.CreateDiaryInvitation(diary)
	if err != nil {
		return nil, err
	}
	return diaryInvitation, nil
}

// VerifyDiaryInvitationCode find diary from invitationCode
func (db *DB) VerifyDiaryInvitationCode(invitationCode string, user User) (*Diary, error) {
	diary := &Diary{}
	if err := db.Joins("JOIN diary_invitations on diaries.id = diary_invitations.diary_id").
		Find(&diary, "diary_invitations.invitation_code = ?", invitationCode).Error; err != nil {
		return nil, err
	}

	if err := db.Create(&DiarySubscriber{UserID: user.ID, DiaryID: diary.ID}).Error; err != nil {
		return nil, err
	}
	return diary, nil
}
