package main

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/suzan2go/familog-api/handler"
	"github.com/suzan2go/familog-api/model"
	"net/http"
)

// ApplicationContext Context for this App
type ApplicationContext struct {
	DatabaseConnection *gorm.DB
	echo.Context
}

// AppError Error struct
type AppError struct {
	code int
	msg  string
}

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
	e := echo.New()
	e.Debug = true
	db := model.InitDB()
	// middleware setting
	e.HTTPErrorHandler = JSONErrorHandler
	e.Use(middleware.Logger())
	h := &handler.Handler{DB: db}

	// routing
	e.GET("/", h.DiaryIndex)

	e.Logger.Fatal(e.Start(":1323"))
}
