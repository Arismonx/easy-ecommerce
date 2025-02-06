package models

import (
	"time"

	"gorm.io/gorm"
)

type Orders struct {
	gorm.Model
	orderDate    time.Time `json:"orderDate"`
	orderAddress string    `json:"orderAddress"`
	orderStatus  string    `json:"orderStatus"`
	orderAmount  float64   `json:"orderAmount"`
	UsersID      uint
	Users        Users `gorm:"foreignKey:"UsersID"`
}
