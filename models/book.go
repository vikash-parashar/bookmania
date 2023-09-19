package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title   string   `json:"book_title"`
	Authors []Author `json:"book_authors" gorm:"many2many:book_authors;"`
	Price   string   `json:"book_price"`
	Year    string   `json:"book_publish_year"`
}
