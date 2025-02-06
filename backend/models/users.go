package models

import (
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Email     string `json:"email" gorm:"unique"`
	Password  string
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
}
