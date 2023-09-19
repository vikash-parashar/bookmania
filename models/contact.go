package models

import "gorm.io/gorm"

type Contact struct {
	gorm.Model
	Name      string  `json:"name"`
	Email     string  `json:"email" gorm:"unique"`
	Phone     string  `json:"phone"`
	Website   string  `json:"website"`
	AddressID uint    `json:"-"`
	Address   Address `json:"address" gorm:"foreignKey:AddressID"`
}
