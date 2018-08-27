package helpers

import (
	"gopkg.in/go-playground/validator.v9"
)

// JSONErrors struct
type JSONErrors struct {
	Field string `json:"field"`
	Rule  string `json:"rule"`
}

// Validate Helper
func Validate(s interface{}) []JSONErrors {
	validate := validator.New()
	err := validate.Struct(s)
	if err != nil {
		errors := []JSONErrors{}

		///
		if _, ok := err.(*validator.InvalidValidationError); ok {
			println(err)
			return nil
		}
		///

		for _, err := range err.(validator.ValidationErrors) {
			n := JSONErrors{Field: err.Field(), Rule: err.ActualTag()}
			errors = append(errors, n)
		}
		return errors
	}
	return nil
}
