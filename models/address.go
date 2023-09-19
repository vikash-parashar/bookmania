package models

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	HNo     string `json:"house_no"`
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	PinCode string `json:"pin_code"`
}
