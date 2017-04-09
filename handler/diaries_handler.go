package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/suzan2go/familog-api/model"
)

// DiaryIndex return DiaryIndex json
func (h Handler) DiaryIndex(c echo.Context) error {
	diary := model.Diary{}
	if err := h.DB.First(&diary, 1).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, diary)
}
