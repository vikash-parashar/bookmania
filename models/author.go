package models

import "gorm.io/gorm"

type Author struct {
	gorm.Model
	Name    string `json:"author_name"`
	Email   string `json:"email" gorm:"unique"`
	Phone   string `json:"phone"`
	Website string `json:"website"`
	HNo     string `json:"house_no"`
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	PinCode string `json:"pin_code"`
	Books   []Book `gorm:"many2many:book_authors;" json:"books"`
}
