package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/stretchr/testify/assert"
	"github.com/suusan2go/familog-api/domain/model"
	"github.com/suusan2go/familog-api/lib/token_generator"
	"github.com/suusan2go/familog-api/registry"
)

// PostSession return DiaryIndex json
func TestPostSession(t *testing.T) {
	db, _ := model.InitTestDB(t)
	deviceToken := tokenGenerator.GenerateRandomToken(32)
	db.FindOrCreateDeviceByToken(deviceToken)

	e := echo.New()
	e.Debug = true
	// middleware setting
	e.Use(middleware.Logger())
	// set form value
	requestBody, _ := json.Marshal(&sessionHandlerPostRequest{DeviceToken: deviceToken})
	req, err := http.NewRequest(
		echo.POST,
		"/session",
		strings.NewReader(string(requestBody)),
	)
	if assert.NoError(t, err) {
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		log.Println(c.FormParams())
		h := &Handler{DB: &db, Registry: &registry.Registry{DB: &db}}

		// Assertions
		if assert.NoError(t, h.PostSession(c)) {
			assert.Equal(t, http.StatusCreated, rec.Code)
		}
	}
}
