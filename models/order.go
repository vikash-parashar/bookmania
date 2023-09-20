package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID uint    `json:"user_id"`
	User   User    `gorm:"foreignKey:UserID" json:"-"`
	Books  []Book  `gorm:"many2many:order_books;" json:"order_books_list"`
	Total  float64 `json:"order_price"`
}
