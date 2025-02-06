package models

import (
	"gorm.io/gorm"
)

type Products struct {
	gorm.Model
	ProductName        string  `json:"productName" gorm:"unique"`
	ProductDescription string  `json:"productDescription"`
	ProductQuantity    float64 `json:"productQuantity"`
	ProductPrice       float64 `json:"productPrice"`
	ProductUom         string  `json:"productUom"`
}
