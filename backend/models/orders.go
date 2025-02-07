package models

import (
	"time"

	"gorm.io/gorm"
)

type Orders struct {
	gorm.Model
	OrderDate    time.Time `json:"orderDate"`
	OrderAddress string    `json:"orderAddress"`
	OrderStatus  string    `json:"orderStatus"`
	OrderAmount  float64   `json:"orderAmount"`
	UserID       uint
	User         Users
}
