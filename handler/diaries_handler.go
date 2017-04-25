package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

// DiaryIndex return DiaryIndex json
func (h *Handler) DiaryIndex(c echo.Context) error {
	ac := c.(*AuthenticatedContext)
	diaries, err := h.DB.AllDiaries(&ac.CurrentUser)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, diaries)
}

// PostDiary Create diary
func (h *Handler) PostDiary(c echo.Context) error {
	ac := c.(*AuthenticatedContext)
	diary, err := h.DB.CreateDiary(&ac.CurrentUser, c.FormValue("title"))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, diary)
}
