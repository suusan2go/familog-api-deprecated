package handler

import (
	"github.com/labstack/echo"
	"net/http"
)

type pushNotificationTokenHandlerRequest struct {
	DeviceToken           string
	PushNotificationToken string
}

// PostPushNotificationToken return nothing
func (h *Handler) PostPushNotificationToken(c echo.Context) error {
	d := &pushNotificationTokenHandlerRequest{}
	if err := c.Bind(d); err != nil {
		return err
	}
	device, err := h.DB.SetPushNotificationToken(d.DeviceToken, d.PushNotificationToken)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, device)
}
