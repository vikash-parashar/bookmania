package controllers

import (
	"myapp/models"

	"gorm.io/gorm"
)

// CreateOrder creates a new order in the database.
func CreateOrder(db *gorm.DB, order *models.Order) error {
	result := db.Create(&order)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// GetOrderByID retrieves an order by its ID.
func GetOrderByID(db *gorm.DB, id uint) (*models.Order, error) {
	var order models.Order
	result := db.First(&order, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &order, nil
}

// UpdateOrder updates an existing order in the database.
func UpdateOrder(db *gorm.DB, order *models.Order) error {
	result := db.Save(order)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// DeleteOrder deletes an order by its ID.
func DeleteOrder(db *gorm.DB, id uint) error {
	result := db.Delete(&models.Order{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// GetAllOrders retrieves all orders from the database.
func GetAllOrders(db *gorm.DB) ([]models.Order, error) {
	var orders []models.Order
	result := db.Find(&orders)
	if result.Error != nil {
		return nil, result.Error
	}
	return orders, nil
}
