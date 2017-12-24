package handler

import (
	"bytes"
	"encoding/json"
	"mime/multipart"
	"net/http"

	"github.com/labstack/echo"
	"github.com/suusan2go/familog-api/domain/service"
)

// PostDiaryEntry Create diary_entry
func (h *Handler) PostDiaryEntry(c echo.Context) error {
	ac := c.(*AuthenticatedContext)
	repo := h.Registry.DiaryRepository()
	diary, err := repo.FindDiary(c.Param("id"), &ac.CurrentUser)
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
		&ac.CurrentUser,
		diary,
		c.FormValue("title"),
		c.FormValue("body"),
		c.FormValue("emoji"),
		images,
	)
	if err != nil {
		return err
	}

	if err := service.DiaryEntryNotificationService(h.Registry.DeviceRepository(), diaryEntry); err != nil {
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
	ac := c.(*AuthenticatedContext)
	diaryEntry, err := h.DB.FindMyDiaryEntry(c.Param("id"), &ac.CurrentUser)
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
	if err := h.DB.UpdateDiaryEntry(
		&ac.CurrentUser,
		diaryEntry,
		c.FormValue("title"),
		c.FormValue("body"),
		c.FormValue("emoji"),
		images,
	); err != nil {
		return err
	}
	diaryEntry, _ = h.DB.FindMyDiaryEntry(c.Param("id"), &ac.CurrentUser)
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	_ = enc.Encode(diaryEntry)
	return c.JSON(http.StatusOK, diaryEntry)
}

// GetDiaryEntries diary entries index
func (h *Handler) GetDiaryEntries(c echo.Context) error {
	ac := c.(*AuthenticatedContext)
	repo := h.Registry.DiaryRepository()
	diary, e := repo.FindDiary(c.Param("id"), &ac.CurrentUser)
	if e != nil {
		return e
	}

	if maxID := c.QueryParam("max_id"); len(maxID) != 0 {
		diaryEntries, err := h.DB.MoreNewerDiaryEntries(diary, maxID)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, diaryEntries)
	}
	if sinceID := c.QueryParam("since_id"); len(sinceID) != 0 {
		diaryEntries, err := h.DB.MoreOlderDiaryEntries(diary, sinceID)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, diaryEntries)
	}
	diaryEntries, err := h.DB.AllDiaryEntries(diary)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, diaryEntries)
}

// GetDiaryEntry diary entry
func (h *Handler) GetDiaryEntry(c echo.Context) error {
	ac := c.(*AuthenticatedContext)
	repo := h.Registry.DiaryEntryRepository()
	diaryEntry, e := repo.FindDiaryEntry(&ac.CurrentUser, c.Param("id"))
	if e != nil {
		return e
	}
	return c.JSON(http.StatusOK, diaryEntry)
}
