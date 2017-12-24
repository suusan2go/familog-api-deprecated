package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

type diaryInvitationVerificationsPostRequest struct {
	InvitationCode string
}

// GetDiaryInvitation return DiaryInvitation Json
func (h *Handler) GetDiaryInvitation(c echo.Context) error {
	ac := c.(*AuthenticatedContext)
	repo := h.Registry.DiaryRepository()
	diary, err := repo.FindDiary(c.Param("id"), &ac.CurrentUser)
	if err != nil {
		return err
	}
	diaryInvitation, err := h.DB.FindNotExpiredDiaryInvitation(diary)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, diaryInvitation)
}

// PostDiaryInvitation return DiaryInvitation Json
func (h *Handler) PostDiaryInvitation(c echo.Context) error {
	ac := c.(*AuthenticatedContext)
	repo := h.Registry.DiaryRepository()
	diary, err := repo.FindDiary(c.Param("id"), &ac.CurrentUser)
	if err != nil {
		return err
	}
	diaryInvitation, err := h.DB.RecreateDiaryInvitation(diary)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, diaryInvitation)
}

// PostDiaryInvitationVerification return DiaryInvitation Json
func (h *Handler) PostDiaryInvitationVerification(c echo.Context) error {
	ac := c.(*AuthenticatedContext)
	r := &diaryInvitationVerificationsPostRequest{}
	if err := ac.Bind(r); err != nil {
		return err
	}
	diary, err := h.DB.VerifyDiaryInvitationCode(r.InvitationCode, ac.CurrentUser)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, diary)
}
