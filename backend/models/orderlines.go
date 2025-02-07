package models

import (
	"gorm.io/gorm"
)

type Orderlines struct {
	gorm.Model
	OrderLineDescription   string  `json:"cartQuantity"`
	OrderLineQuantity      float64 `json:"orderLineQuantity"`
	OrderLineUnitPrice     float64 `json:"orderLineUnitPrice"`
	OrderLinePriceSubtotal float64 `json:"orderLinePriceSubtotal"`
	OrderID                uint
	ProductID              uint
	Order                  Orders
	Product                Products
}
