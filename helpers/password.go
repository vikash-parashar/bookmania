package helpers

import "golang.org/x/crypto/bcrypt"

// HashPassword takes a plain text password and returns its bcrypt hash as []byte.
func HashPassword(password string) ([]byte, error) {
	// Generate a salt and hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return hashedPassword, nil
}
