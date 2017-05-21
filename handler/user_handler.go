package handler

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo"
)

// PatchUser Patch user
func (h *Handler) PatchUser(c echo.Context) error {
	ac := c.(*AuthenticatedContext)
	file, _ := c.FormFile("image")
	user := &ac.CurrentUser
	if err := h.DB.UpdateUser(
		user,
		c.FormValue("name"),
		file,
	); err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	_ = enc.Encode(user)
	return c.JSONBlob(http.StatusOK, buf.Bytes())
}

// GetUser current user
func (h *Handler) GetUser(c echo.Context) error {
	ac := c.(*AuthenticatedContext)

	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	_ = enc.Encode(ac.CurrentUser)
	return c.JSONBlob(http.StatusOK, buf.Bytes())
}
