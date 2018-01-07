package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

type sessionHandlerPostRequest struct {
	DeviceToken string
}

// PostSession return DiaryIndex json
func (h *Handler) PostSession(c echo.Context) error {
	repo := h.Registry.SessionRepository()
	r := &sessionHandlerPostRequest{}
	if err := c.Bind(r); err != nil {
		return err
	}
	user, err := h.DB.FindUserByDeviceToken(r.DeviceToken)
	if err != nil {
		return err
	}
	sessionToken, err := repo.GenerateOrExtendSessionToken(user)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, sessionToken)
}
