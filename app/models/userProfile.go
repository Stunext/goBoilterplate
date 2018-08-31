package models

import (
	"goBoilterplate/config"
)

// UserProfile model
type UserProfile struct {
	ID       int     `json:"id" gorm:"primary_key"`
	Username string  `validate:"required" json:"username"`
	Email    string  `validate:"required,email" json:"email"`
	Password string  `validate:"required" json:"password"`
	Role     string  `validate:"required" json:"role"`
	Profile  Profile `json:"profile" gorm:"foreignkey:UserID"`
}

// TableName UserProfile to be `users`
func (UserProfile) TableName() string {
	return "users"
}

// UserProfileList Listado de Usuarios con sus datos de perfil
func UserProfileList() []UserProfile {
	users := []UserProfile{}
	res := config.DB.Debug().Preload("Profile").Find(&users)
	if res.Error == nil {
		return users
	}
	return nil
}

// UserProfileShow Consulta un Usuario con sus datos de perfil
func UserProfileShow(id int) *UserProfile {
	user := new(UserProfile)
	res := config.DB.First(&user, id).Related(&user.Profile, "UserId")
	if res.Error == nil {
		return user
	}
	return nil
}
