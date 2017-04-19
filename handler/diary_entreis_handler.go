package handler

import (
	"bytes"
	"encoding/json"
	"github.com/labstack/echo"
	"mime/multipart"
	"net/http"
)

// PostDiaryEntry Create diary_entry
func (h *Handler) PostDiaryEntry(c echo.Context) error {
	diary, err := h.DB.FindDiary(c.Param("id"), h.CurrentUser)
	if err != nil {
		return err
	}
	file1, _ := c.FormFile("image1")
	file2, _ := c.FormFile("image2")
	file3, _ := c.FormFile("image3")
	images := []*multipart.FileHeader{
		file1,
		file2,
		file3,
	}
	diaryEntry, err := h.DB.CreateDiaryEntry(
		h.CurrentUser,
		diary,
		c.FormValue("title"),
		c.FormValue("body"),
		c.FormValue("emoji"),
		images,
	)

	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	_ = enc.Encode(diaryEntry)
	return c.JSONBlob(http.StatusOK, buf.Bytes())
}

// PatchDiaryEntry Create diary
func (h *Handler) PatchDiaryEntry(c echo.Context) error {
	diaryEntry, err := h.DB.FindMyDiaryEntry(c.Param("id"), h.CurrentUser)
	if err != nil {
		return err
	}
	diaryEntry.Title = c.FormValue("title")
	diaryEntry.Body = c.FormValue("body")
	diaryEntry.Emoji = c.FormValue("emoji")
	if err := h.DB.Save(diaryEntry).Error; err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	_ = enc.Encode(diaryEntry)
	return c.JSON(http.StatusOK, diaryEntry)
}
