package models

import "gorm.io/gorm"

// User struct
type User struct {
	gorm.Model
	Name    string `json:"name"`
	Email   string `json:"email"`
	Website string `json:"website"`
}
