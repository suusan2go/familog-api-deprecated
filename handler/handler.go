package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/suzan2go/familog-api/model"
)

// Handler Handle Http Request
type Handler struct {
	DB *model.DB
}

// AuthenticatedContext include CurrentUser
type AuthenticatedContext struct {
	echo.Context
	CurrentUser model.User
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
		ac := &AuthenticatedContext{c, sessionToken.User}
		return next(ac)
	}
}
