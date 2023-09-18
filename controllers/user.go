package controllers

import (
	"myapp/models"

	"gorm.io/gorm"
)

// GetAllUsers retrieves all users from the database.
func GetAllUsers(db *gorm.DB) ([]models.User, error) {
	var users []models.User
	result := db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

// CreateUser creates a new user in the database.
func CreateUser(db *gorm.DB, user *models.User) error {
	result := db.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// GetUserByID retrieves a user by their ID.
func GetUserByID(db *gorm.DB, id uint) (*models.User, error) {
	var user models.User
	result := db.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// UpdateUser updates an existing user in the database.
func UpdateUser(db *gorm.DB, user *models.User) error {
	result := db.Save(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// DeleteUser deletes a user by their ID.
func DeleteUser(db *gorm.DB, id uint) error {
	result := db.Delete(&models.User{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// SearchUser searches for a user in the database based on email, phone, first name, or last name.
func SearchUser(db *gorm.DB, email, phone, firstName, lastName string) ([]models.User, error) {
	var users []models.User

	// Build the query conditions based on the provided parameters
	query := db.Where("1 = 1") // A dummy condition to build the query dynamically

	if email != "" {
		query = query.Where("email = ?", email)
	}

	if phone != "" {
		query = query.Where("phone = ?", phone)
	}

	if firstName != "" {
		query = query.Where("first_name = ?", firstName)
	}

	if lastName != "" {
		query = query.Where("last_name = ?", lastName)
	}

	// Execute the query and retrieve the matching users
	if err := query.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

// emailExists checks if a user with the given email already exists in the database.
func IsEmailExists(db *gorm.DB, email string) bool {
	var user models.User
	if db.Where("email = ?", email).First(&user).Error == nil {
		return true
	}
	return false
}
