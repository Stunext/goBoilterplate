package models

import (
	"encoding/json"
	"goBoilterplate/config"
	"strconv"

	"github.com/dgrijalva/jwt-go"
)

// JwtClaims struct
type JwtClaims struct {
	ID int `json:"id"`
	jwt.StandardClaims
}

// Login struct
type Login struct {
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required" json:"password"`
}

// AuthLogin Login Query
func AuthLogin(email string, password string) *User {
	user := new(User)
	res := config.DB.Where("email = ? and password = ?", email, password).First(&user)
	if res.Error == nil {

		json, err := json.Marshal(&user)
		if err == nil {
			config.RC.Set(strconv.Itoa(user.ID), json, 0).Err()
		}

		return user
	}
	return nil
}
