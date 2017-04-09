package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

// DiaryIndex return DiaryIndex json
func (h Handler) DiaryIndex(c echo.Context) error {
	diaries, err := h.DB.AllDiaries()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, diaries)
}
