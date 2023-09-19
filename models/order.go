package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID    uint    `json:"user_id"`
	User      User    `json:"user" gorm:"foreignKey:UserID"`
	Books     []Book  `json:"order_books_list" gorm:"many2many:order_books;"`
	Total     float64 `json:"order_price"`
	AddressID uint    `json:"-"`
	ContactID uint    `json:"-"`
	Address   Address `json:"order_address" gorm:"foreignKey:AddressID"`
	Contact   Contact `json:"order_contact_details" gorm:"foreignKey:ContactID"`
}
