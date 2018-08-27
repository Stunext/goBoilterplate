package tests

import (
	"encoding/json"
	"net/http/httptest"
	"net/url"
	"goBoilterplate/app/controllers"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestHome(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	if assert.NoError(t, controllers.Index(c)) {
		assert.Equal(t, 200, rec.Code)
	}
}

type Login struct {
	Token string
}

var JWT *Login

func TestLogin(t *testing.T) {
	e := echo.New()

	// User found

	f := make(url.Values)
	f.Set("email", "andres@teachlr.org")
	f.Set("password", "123456")

	req := httptest.NewRequest(echo.POST, "/login", strings.NewReader(f.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	if assert.NoError(t, controllers.Login(c)) {
		if err := json.NewDecoder(rec.Body).Decode(&JWT); err != nil {
			panic(err)
		}
		req.Body.Close()

		assert.Equal(t, 200, rec.Code)
	}

	// User not found

	f = make(url.Values)
	f.Set("email", "andres@teachlr.org")
	f.Set("password", "1234567")

	req = httptest.NewRequest(echo.POST, "/login", strings.NewReader(f.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec = httptest.NewRecorder()

	c = e.NewContext(req, rec)

	if assert.NoError(t, controllers.Login(c)) {
		assert.Equal(t, 404, rec.Code)
	}
}
