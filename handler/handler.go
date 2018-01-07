package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/suusan2go/familog-api/domain/model"
	"github.com/suusan2go/familog-api/registry"
)

// Handler Handle Http Request
type Handler struct {
	DB       *model.DB
	Registry *registry.Registry
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
		repo := h.Registry.SessionRepository()
		sessionToken, err := repo.FindSessionToken(token)
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

// GetAppInfo return AppInfo for healthcheck
func (h *Handler) GetAppInfo(c echo.Context) error {
	type AppInfo struct {
		Status string `json:"status"`
	}
	return c.JSON(http.StatusOK, &AppInfo{Status: "ok"})
}
