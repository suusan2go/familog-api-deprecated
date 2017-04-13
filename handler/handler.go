package handler

import (
	"github.com/labstack/echo"
	"github.com/suzan2go/familog-api/model"
	"net/http"
)

// Handler Handle Http Request
type Handler struct {
	DB          *model.DB
	CurrentUser *model.User
}

// Authenticate and Set CurrentUser
func (h *Handler) Authenticate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")
		sessionToken, err := h.DB.FindSessionToken(token)
		if err != nil {
			return echo.NewHTTPError(http.StatusForbidden, err)
		}
		if !sessionToken.IsValid() {
			return echo.NewHTTPError(http.StatusForbidden, "This token is expired")
		}
		h.CurrentUser = &(sessionToken.User)
		return next(c)
	}
}
