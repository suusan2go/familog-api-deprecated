package handler

import (
	"github.com/labstack/echo"
	"log"
	"net/http"
)

// PostDevice return Device json
func (h *Handler) PostDevice(c echo.Context) error {
	log.Printf(c.FormValue("deviceToken"))
	device, err := h.DB.FindOrCreateDeviceByToken(c.FormValue("deviceToken"))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, device)
}
