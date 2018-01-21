package handler

import (
	"bytes"
	"encoding/json"
	"mime/multipart"
	"net/http"

	"github.com/labstack/echo"
	"github.com/suusan2go/familog-api/domain/service"
	"github.com/suusan2go/familog-api/domain/model"
)

// PostDiaryEntry Create diary_entry
func (h *Handler) PostDiaryEntry(c echo.Context) error {
	ac := c.(*AuthenticatedContext)
	diary, err := h.Registry.DiaryRepository().FindDiary(c.Param("id"), &ac.CurrentUser)
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
	diaryEntry := &model.DiaryEntry{
		Title: c.FormValue("title"),
		Body: c.FormValue("body"),
		Emoji: c.FormValue("emoji"),
		DiaryID: diary.ID,
		UserID: ac.CurrentUser.ID,
	}
	if err := h.Registry.DiaryEntryRepository().Save(diaryEntry); err != nil {
		return err
	}
	for _, image := range images {
		if image == nil {
			continue
		}
		diaryEntryImage := model.MapImageToDiaryEntryImage(image, *diaryEntry)
		if err := h.Registry.DiaryEntryImageRepository().Save(diaryEntryImage); err != nil {
			return err
		}
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
	repo := h.Registry.DiaryEntryRepository()
	diaryEntry, err := repo.FindMyDiaryEntry(c.Param("id"), &ac.CurrentUser)
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
	diaryEntry, _ = repo.FindMyDiaryEntry(c.Param("id"), &ac.CurrentUser)
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
	repoDe := h.Registry.DiaryEntryRepository()
	diaryEntries, err := repoDe.AllDiaryEntries(diary)
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
