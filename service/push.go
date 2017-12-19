package service

import (
	"io/ioutil"
	"log"

	"github.com/suusan2go/familog-api/domain/model"
	"github.com/suusan2go/familog-api/lib/push"
)

// PushNotificationToSubscriver push notification to user
func PushNotificationToSubscriver(db *model.DB, d *model.DiaryEntry) error {
	var devices []*model.Device
	var payloads []*push.NotificationPayload
	if err := db.Where("diary_subscribers.diary_id =?", d.DiaryID).
		Where("diary_subscribers.user_id != ?", d.UserID).
		Joins("JOIN users ON users.id = devices.user_id").
		Joins("JOIN diary_subscribers ON diary_subscribers.user_id = users.id").
		Find(&devices).Error; err != nil {
		return err
	}
	for _, dv := range devices {
		payloads = append(payloads,
			&push.NotificationPayload{
				To: dv.PushNotificationToken, Body: d.Title, Badge: 1,
			},
		)
	}
	if response, err := push.Push(payloads); err != nil {
		log.Print(err)
		b, _ := ioutil.ReadAll(response.Body)
		log.Println(string(b))
	}

	return nil
}
