package main

import (
	"myapp/config"
	"myapp/handlers"
	"net/http"
)

func init() {
	config.SetupDB()

}
func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/login", handlers.LogIn)
	http.HandleFunc("/logout", handlers.LogOut)

	// routes for user
	http.HandleFunc("/user/all", handlers.GetAllUsers)
	http.HandleFunc("/user/create", handlers.CreateUser)

	http.HandleFunc("/user/profile", handlers.GetUserProfile)
	http.HandleFunc("/user/get/id", handlers.GetUserByID)
	http.HandleFunc("/user/update/id", handlers.UpdateUser)
	http.HandleFunc("/user/delete/id", handlers.DeleteUser)

	// routes for book
	http.HandleFunc("/book/all", handlers.GetAllBooks)
	http.HandleFunc("/book/create", handlers.CreateBook)
	http.HandleFunc("/book/get/id", handlers.GetBookByID)
	http.HandleFunc("/book/update/id", handlers.UpdateBook)
	http.HandleFunc("/book/delete/id", handlers.DeleteBook)

	// routes for order
	http.HandleFunc("/order/all", handlers.GetAllOrders)
	http.HandleFunc("/order/create", handlers.CreateOrder)
	http.HandleFunc("/order/get/id", handlers.GetOrderByID)
	http.HandleFunc("/order/update/id", handlers.UpdateOrder)
	http.HandleFunc("/order/delete/id", handlers.DeleteOrder)

	http.ListenAndServe(":8080", nil)
}
