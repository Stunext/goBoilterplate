package models

import (
	"goBoilterplate/config"
)

// User model
type User struct {
	ID       int      `json:"id" gorm:"primary_key"`
	Username string   `validate:"required" json:"username"`
	Email    string   `validate:"required,email" json:"email"`
	Password string   `validate:"required" json:"password"`
	Role     string   `validate:"required" json:"role"`
}

// UserList Listado de Usuarios
func UserList() []*User {
	users := []*User{}
	res := config.DB.Find(&users)
	if res.Error == nil {
		return users
	}
	return nil
}

// UserStore Almacena un Usuario
func UserStore(user *User) bool {
	res := config.DB.Create(&user)
	if res.Error == nil {
		return true
	}
	return false
}

// UserShow Consulta un Usuario
func UserShow(id int) *User {
	user := new(User)
	res := config.DB.First(&user, id)
	if res.Error == nil {
		return user
	}
	return nil
}

// UserUpdate Actualiza un Usuario
func UserUpdate(user *User) bool {
	res := config.DB.Save(&user)
	if res.Error == nil {
		return true
	}
	return false
}

// UserDelete Elimina un Usuario
func UserDelete(id int) bool {
	user := new(User)
	res := config.DB.Where("id = ?", id).Delete(&user)
	if res.Error == nil {
		return true
	}
	return false
}
