package handler

import (
	"github.com/labstack/echo"
	"net/http"
)

type deviceHandlerRequest struct {
	DeviceToken string
}

// PostDevice return Device json
func (h *Handler) PostDevice(c echo.Context) error {
	d := &deviceHandlerRequest{}
	if err := c.Bind(d); err != nil {
		return err
	}
	device, err := h.DB.FindOrCreateDeviceByToken(d.DeviceToken)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, device)
}
