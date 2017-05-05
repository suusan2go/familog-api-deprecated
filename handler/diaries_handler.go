package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

type diariesHandlerPostRequest struct {
	Title string
}

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
	r := &diariesHandlerPostRequest{}
	if err := c.Bind(r); err != nil {
		return err
	}
	diary, err := h.DB.CreateDiary(&ac.CurrentUser, r.Title)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, diary)
}
