package service

import (
	"io/ioutil"
	"log"

	"github.com/suusan2go/familog-api/domain/model"
	"github.com/suusan2go/familog-api/domain/repository"
	"github.com/suusan2go/familog-api/infrastructure/expo/push"
)

// DiaryEntryNotificationService notification to user
// TODO: create struct for this service and use constructor injection
func DiaryEntryNotificationService(repo repository.DeviceRepository, d *model.DiaryEntry) error {
	devices, err := repo.FindSubscribers(d)
	if err != nil {
		return err
	}
	// TODO: move thease logic to infrastracutre layer
	var payloads []*push.NotificationPayload
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
