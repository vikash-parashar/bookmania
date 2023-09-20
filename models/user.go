package models

import "gorm.io/gorm"

const (
	Admin   = "admin"
	General = "General"
)

type User struct {
	gorm.Model
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Email     string  `json:"email" gorm:"unique"`
	Phone     string  `json:"phone"`
	Orders    []Order `json:"orders"`
	Password  []byte  `json:"password"`
	UserType  string  `json:"user_type"`
	HNo       string  `json:"house_no"`
	Street    string  `json:"street"`
	City      string  `json:"city"`
	State     string  `json:"state"`
	PinCode   string  `json:"pin_code"`
}
