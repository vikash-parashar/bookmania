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
	AddressID uint    `json:"-"`
	ContactID uint    `json:"-"`
	Orders    []Order `json:"orders"`
	Password  []byte  `json:"password"`
	UserType  string  `json:"user_type"`
	Address   Address `json:"address" gorm:"foreignKey:AddressID"`
	Contact   Contact `json:"contact" gorm:"foreignKey:ContactID"`
}
