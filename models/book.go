package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	// Id          int    `json:"book_id"`
	Title       string `json:"book_title"`
	FirstName   string `json:"author_first_name"`
	LastName    string `json:"author_last_name"`
	LinkedIn    string `json:"linked_in"`
	Facebook    string `json:"facebook"`
	Website     string `json:"website"`
	Email       string `json:"contact_email"`
	Phone       string `json:"contact_phone"`
	PublishYear int    `json:"book_publish_year"`
	Price       int    `json:"book_price"`
}
