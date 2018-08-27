package models

// Profile model
type Profile struct {
	ID        int    `json:"-" gorm:"primary_key"`
	UserID    int    `validate:"required" json:"user_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Gender    string `json:"gender"`
	Age       int    `json:"age"`
}
