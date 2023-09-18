package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	Quantity        int     `json:"order_quantity"`
	TotalOrderPrice float64 `json:"total_order_price"`
}
