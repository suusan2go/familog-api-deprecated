package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

// PostSession return DiaryIndex json
func (h Handler) PostSession(c echo.Context) error {
	user, err := h.DB.FindUserByDeviceToken(c.FormValue("deviceToken"))
	if err != nil {
		return err
	}
	sessionToken, err := h.DB.GenerateOrExtendSessionToken(user)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, sessionToken)
}
