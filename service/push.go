package service

import (
	"log"

	"github.com/suzan2go/familog-api/lib/push"
	"github.com/suzan2go/familog-api/model"
)

// PushNotificationToSubscriver push notification to user
func PushNotificationToSubscriver(db *model.DB, d *model.DiaryEntry) error {
	var devices []*model.Device
	if err := db.Where("diary_subscribers.diary_id =?", d.DiaryID).
		Joins("JOIN users ON users.id = devices.user_id").
		Joins("JOIN diary_subscribers ON diary_subscribers.user_id = users.id").
		Find(&devices).Error; err != nil {
		return err
	}
	for _, dv := range devices {
		n := &push.NotificationPayload{To: dv.PushNotificationToken, Title: d.Title, Badge: 1}
		if response, err := push.Push(n); err != nil {
			log.Print(err)
			log.Print(response)
		}
	}
	return nil
}
