package models

import (
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	CartQuantity float64 `json:"cartQuantity"`
	ProductID    uint
	UserID       uint
	Product      Products
	User         Users
}
