package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title   string   `json:"book_title"`
	Authors []Author `gorm:"many2many:book_authors;" json:"book_authors"`
	Price   float64  `json:"book_price"`
	Year    string   `json:"book_publish_year"`
	Orders  []Order  `gorm:"many2many:order_books;" json:"-"`
}
