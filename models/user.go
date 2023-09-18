package models

import "gorm.io/gorm"

const (
	General = "general"
	Admin   = "Admin"
)

type User struct {
	gorm.Model
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	BuildingNumber string `json:"building_number"`
	City           string `json:"city"`
	State          string `json:"state"`
	Country        string `json:"country"`
	PinCode        string `json:"pin_code"`
	Email          string `json:"email"`
	Phone          string `json:"phone"`
	Password       []byte `json:"password"`
	Type           string `json:"user_type"`
}
