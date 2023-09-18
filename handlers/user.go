package handlers

import (
	"encoding/json"
	"log"
	"myapp/config"
	"myapp/controllers"
	"myapp/helpers"
	"myapp/models"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := config.DB

	// Parse form data from the request
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}

	// Extract email from the form data
	email := r.FormValue("email")

	// Check if a user with the same email already exists
	if controllers.IsEmailExists(db, email) {
		errorResponse := map[string]interface{}{
			"status":  "error",
			"message": "User with this email already exists",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusConflict)
		if err = json.NewEncoder(w).Encode(errorResponse); err != nil {
			log.Println(err)
		}
		return
	}

	// Create a User struct and populate it with form values
	user := &models.User{
		FirstName:      r.FormValue("first_name"),
		LastName:       r.FormValue("last_name"),
		BuildingNumber: r.FormValue("building_number"),
		City:           r.FormValue("city"),
		State:          r.FormValue("state"),
		Country:        r.FormValue("country"),
		PinCode:        r.FormValue("pin_code"),
		Email:          email,
		Phone:          r.FormValue("phone"),
		Password:       []byte(r.FormValue("password")),
		Type:           r.FormValue("user_type"),
	}

	// Call the controllers.CreateUser function to create the user
	if err = controllers.CreateUser(db, user); err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		log.Printf("Error creating user: %v", err)
		return
	}
	// Send a welcome email to the user
	if err = helpers.SendWelcomeEmail(user.Email); err != nil {
		log.Printf("error sending welcome email: %v", err)
		// Handle the email sending error if needed
	}
	log.Printf("welcome email sent successfully to : %s\n", user.Email)
	// Respond with a success message
	response := map[string]interface{}{
		"status":  "success",
		"message": "new user is created successfully",
		"data":    *user, // Dereference the user pointer
	}

	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(response); err != nil {
		log.Println(err)
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

}
func DeleteUser(w http.ResponseWriter, r *http.Request) {

}
func Home(w http.ResponseWriter, r *http.Request) {

}

func LogIn(w http.ResponseWriter, r *http.Request) {

}
func LogOut(w http.ResponseWriter, r *http.Request) {

}
func GetUserProfile(w http.ResponseWriter, r *http.Request) {

}
func GetAllUsers(w http.ResponseWriter, r *http.Request) {

}
func GetUserByID(w http.ResponseWriter, r *http.Request) {

}
