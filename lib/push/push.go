package push

import (
	"bytes"
	"encoding/json"
	"net/http"
)

const pushEndpointURL = "https://exp.host/--/api/v2/push"

// NotificationPayload notification parametes
type NotificationPayload struct {
	To    string `json:"to"`
	Title string `json:"title"`
	Body  string `json:"body"`
	Badge int    `json:"badge"`
}

// Push notification to specific user
func Push(p *NotificationPayload) (*http.Response, error) {
	payload, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	req, _ := http.NewRequest("POST", pushEndpointURL, bytes.NewBuffer(payload))
	req.Header.Set("accept", "application/json")
	req.Header.Set("content-type", "application/json")
	req.Header.Set("accept-encoding", "gzip, deflate")

	client := new(http.Client)
	return client.Do(req)
}
