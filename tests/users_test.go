package tests

import (
	"encoding/json"
	"goBoilterplate/app/controllers"
	"goBoilterplate/app/models"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var User *models.User

func TestUserList(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(echo.GET, "/api/users", nil)
	req.Header.Set(echo.HeaderAuthorization, "Bearer "+JWT.Token)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	if assert.NoError(t, controllers.UserList(c)) {
		assert.Equal(t, 200, rec.Code)
	}
}

func TestUserStore(t *testing.T) {
	e := echo.New()

	// New User

	f := make(url.Values)
	f.Set("name", "Andres Fuentes")
	f.Set("email", "andresf@manzanares.com.ve")
	f.Set("password", "123456")
	f.Set("role", "Admin")

	req := httptest.NewRequest(echo.POST, "/api/users", strings.NewReader(f.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	req.Header.Set(echo.HeaderAuthorization, "Bearer "+JWT.Token)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	if assert.NoError(t, controllers.UserStore(c)) {

		if err := json.NewDecoder(rec.Body).Decode(&User); err != nil {
			panic(err)
		}
		req.Body.Close()

		assert.Equal(t, 201, rec.Code)
	}

	// Validation Error

	f = make(url.Values)
	f.Set("name", "Andres Fuentes")
	f.Set("password", "123456")
	f.Set("role", "Admin")

	req = httptest.NewRequest(echo.POST, "/api/users", strings.NewReader(f.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	req.Header.Set(echo.HeaderAuthorization, "Bearer "+JWT.Token)
	rec = httptest.NewRecorder()

	c = e.NewContext(req, rec)

	if assert.NoError(t, controllers.UserStore(c)) {
		assert.Equal(t, 422, rec.Code)
	}
}

func TestUserShow(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(echo.GET, "/api/users", nil)
	req.Header.Set(echo.HeaderAuthorization, "Bearer "+JWT.Token)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(User.ID))

	if assert.NoError(t, controllers.UserShow(c)) {
		assert.Equal(t, 200, rec.Code)
	}
}

func TestUserUpdate(t *testing.T) {
	e := echo.New()

	// Update User

	f := make(url.Values)
	f.Set("name", "Andres Fuentes A")
	f.Set("email", "andresf@manzanares.com.ve")
	f.Set("password", "123456")
	f.Set("role", "Admin")

	req := httptest.NewRequest(echo.PUT, "/api/users", strings.NewReader(f.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	req.Header.Set(echo.HeaderAuthorization, "Bearer "+JWT.Token)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(User.ID))

	if assert.NoError(t, controllers.UserUpdate(c)) {
		assert.Equal(t, 200, rec.Code)
	}

	// Validation Error

	f = make(url.Values)
	f.Set("name", "Andres Fuentes")
	f.Set("password", "123456")
	f.Set("role", "Admin")

	req = httptest.NewRequest(echo.POST, "/api/users", strings.NewReader(f.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	req.Header.Set(echo.HeaderAuthorization, "Bearer "+JWT.Token)
	rec = httptest.NewRecorder()

	c = e.NewContext(req, rec)

	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(User.ID))

	if assert.NoError(t, controllers.UserUpdate(c)) {
		assert.Equal(t, 422, rec.Code)
	}
}

func TestUserDelete(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(echo.GET, "/api/users/", nil)
	req.Header.Set(echo.HeaderAuthorization, "Bearer "+JWT.Token)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(User.ID))

	if assert.NoError(t, controllers.UserDelete(c)) {
		assert.Equal(t, 200, rec.Code)
	}
}
