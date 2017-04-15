package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

// DiaryIndex return DiaryIndex json
func (h *Handler) DiaryIndex(c echo.Context) error {
	diaries, err := h.DB.AllDiaries(h.CurrentUser)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, diaries)
}

// PostDiary Create diary
func (h *Handler) PostDiary(c echo.Context) error {
	diary, err := h.DB.CreateDiary(h.CurrentUser, c.FormValue("title"))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, diary)
}
