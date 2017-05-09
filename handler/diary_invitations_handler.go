package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

// GetDiaryInvitation return DiaryInvitation Json
func (h *Handler) GetDiaryInvitation(c echo.Context) error {
	ac := c.(*AuthenticatedContext)
	diary, err := h.DB.FindDiary(c.Param("id"), &ac.CurrentUser)
	if err != nil {
		return err
	}
	diaryInvitation, err := h.DB.FindNotExpiredDiaryInvitation(diary)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, diaryInvitation)
}

// CreateDiaryInvitation return DiaryInvitation Json
func (h *Handler) CreateDiaryInvitation(c echo.Context) error {
	ac := c.(*AuthenticatedContext)
	diary, err := h.DB.FindDiary(c.Param("id"), &ac.CurrentUser)
	if err != nil {
		return err
	}
	diaryInvitation, err := h.DB.RecreateDiaryInvitation(diary)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, diaryInvitation)
}
