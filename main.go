package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/suzan2go/familog-api/handler"
	"github.com/suzan2go/familog-api/model"
	"net/http"
)

// Map Generic Map
type Map map[string]interface{}

// JSONErrorHandler Handling Errors as JSON
func JSONErrorHandler(err error, c echo.Context) {
	var (
		code = http.StatusInternalServerError
		msg  interface{}
	)

	e := c.Echo()

	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		msg = he.Message
	} else if e.Debug {
		msg = err.Error()
	} else {
		msg = http.StatusText(code)
	}
	if _, ok := msg.(string); ok {
		msg = Map{"message": msg}
	}

	if !c.Response().Committed {
		if c.Request().Method == "HEAD" {
			if err := c.NoContent(code); err != nil {
				goto ERROR
			}
		} else {
			if err := c.JSON(code, msg); err != nil {
				goto ERROR
			}
		}
	}
ERROR:
	e.Logger.Error(err)
}

func main() {
	db := model.InitDB()
	db.Migration()
	e := echo.New()
	e.Debug = true
	e.HTTPErrorHandler = JSONErrorHandler
	// middleware setting
	e.Use(middleware.Logger())
	h := &handler.Handler{DB: db}

	// routing
	e.POST("/device", h.PostDevice)
	e.POST("/session", h.PostSession)
	e.POST("/diaries", h.PostDiary, h.Authenticate)
	e.GET("/diaries", h.DiaryIndex, h.Authenticate)
	e.POST("/diaries/:id/diary_entries", h.PostDiaryEntry, h.Authenticate)
	e.PATCH("/diary_entries/:id", h.PatchDiaryEntry, h.Authenticate)
	e.PATCH("/diary_entries/:diary_entry_id/images/:id", h.PatchDiaryEntryImage, h.Authenticate)
	e.DELETE("/diary_entries/:diary_entry_id/images/:id", h.DeleteDiaryEntryImage, h.Authenticate)

	e.Logger.Fatal(e.Start(":8080"))
}