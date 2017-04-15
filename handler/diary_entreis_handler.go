package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

// PostDiaryEntry Create diary_entry
func (h *Handler) PostDiaryEntry(c echo.Context) error {
	diary, err := h.DB.FindDiary(c.Param("id"), h.CurrentUser)
	if err != nil {
		return err
	}
	diaryEntry, err := h.DB.CreateDiaryEntry(
		h.CurrentUser,
		diary,
		c.FormValue("title"),
		c.FormValue("body"),
		c.FormValue("emoji"),
	)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, diaryEntry)
}
