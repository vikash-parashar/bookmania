package models

import "gorm.io/gorm"

type Author struct {
	gorm.Model
	Name      string  `json:"author_name"`
	Email     string  `json:"email" gorm:"unique"`
	ContactID uint    `json:"-"`
	Contact   Contact `json:"author_contact" gorm:"foreignKey:ContactID"`
	Books     []Book  `json:"books" gorm:"many2many:book_authors;"`
}
