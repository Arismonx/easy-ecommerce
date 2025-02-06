package models

import (
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	CartQuantity float64 `json:"cartQuantity"`
	ProductID    uint
	UserID       uint
	Products     Products `gorm:"foreignKey:"ProductID"`
	Users        Users    `gorm:"foreignKey:"UserID"`
}
