package controllers

import (
	"goBoilterplate/app/helpers"
	"goBoilterplate/app/models"
	"strconv"

	"github.com/labstack/echo"
)

// UserList godoc
// @Summary UserList
// @Description Listado de Usuarios
// @Tags User
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {array} models.User
// @Router /api/users [get]
func UserList(c echo.Context) error {
	users := models.UserList()
	if users != nil {
		return c.JSON(200, users)
	}
	return c.JSON(204, "No Content")
}

// UserStore godoc
// @Summary UserStore
// @Description Guardar datos de Usuario
// @Tags User
// @Produce json
// @Security ApiKeyAuth
// @Param username query string true "Username"
// @Param email query string true "Email"
// @Param role query string true "Role"
// @Param password query string true "Password"
// @Success 201 {object} models.User
// @Failure 422 {string} string
// @Failure 400 {string} string
// @Router /api/users [post]
func UserStore(c echo.Context) error {
	user := models.User{}
	user.Username = c.FormValue("username")
	user.Email = c.FormValue("email")
	user.Password = c.FormValue("password")
	user.Role = c.FormValue("role")

	err := helpers.Validate(&user)
	if err != nil {
		return c.JSON(422, err)
	}

	res := models.UserStore(&user)
	if res {
		return c.JSON(201, user)
	}
	return c.JSON(400, "Bad Request")
}

// UserShow godoc
// @Summary UserShow
// @Description Consultar Usuario
// @Tags User
// @Produce  json
// @Security ApiKeyAuth
// @Param id path int true "Id"
// @Success 200 {object} models.User
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Router /api/users/{id} [get]
func UserShow(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err == nil {
		user := models.UserProfileShow(id)
		if user != nil {
			return c.JSON(200, user)
		}
		return c.JSON(404, "Not Found")
	}
	return c.JSON(400, "Bad Request")
}

// UserUpdate godoc
// @Summary UserUpdate
// @Description Actualizar datos de Usuario
// @Tags User
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Id"
// @Param username query string true "Username"
// @Param email query string true "Email"
// @Param role query string true "Role"
// @Param password query string true "Password"
// @Success 200 {object} models.User
// @Failure 422 {string} string
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Router /api/users/{id} [put]
func UserUpdate(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err == nil {
		user := models.UserShow(id)
		if user != nil {
			user.Username = c.FormValue("username")
			user.Email = c.FormValue("email")
			user.Password = c.FormValue("password")
			user.Role = c.FormValue("role")

			err := helpers.Validate(user)
			if err != nil {
				return c.JSON(422, err)
			}

			res := models.UserUpdate(user)
			if res {
				return c.JSON(200, user)
			}
		} else {
			return c.JSON(404, "Not Found")
		}
	}
	return c.JSON(400, "Bad Request")
}

// UserDelete godoc
// @Summary UserDelete
// @Description Borrado de Usuario
// @Tags User
// @Produce  json
// @Security ApiKeyAuth
// @Param id path string true "Id"
// @Success 200 {string} string
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Router /api/users/{id} [delete]
func UserDelete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err == nil {
		res := models.UserDelete(id)
		if res {
			return c.JSON(200, "Deleted")
		}
		return c.JSON(404, "Not Found")
	}
	return c.JSON(400, "Bad Request")
}
