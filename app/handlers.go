package app

import (
	"boilerplate/controllers"

	"github.com/gorilla/mux"
)

// LoadRoutes is a function to create and return a new *mux.Router and your routes.
func LoadRoutes() *mux.Router {
	router := mux.NewRouter()

	// User handlers
	router.HandleFunc("/api/users", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/api/users", controllers.GetAllUsers).Methods("GET")
	router.HandleFunc("/api/users/{id}", controllers.GetUserByID).Methods("GET")
	router.HandleFunc("/api/users", controllers.UpdateUserByID).Methods("PUT")
	router.HandleFunc("/api/users/{id}", controllers.DeleteUserByID).Methods("DELETE")
	router.HandleFunc("/api/users/{id}/creditcards", controllers.GetAllCreditCardsByUserID).Methods("GET")

	// Credit Card handlers
	router.HandleFunc("/api/creditcards", controllers.CreateCreditCard).Methods("POST")
	router.HandleFunc("/api/creditcards", controllers.GetAllCreditCards).Methods("GET")
	router.HandleFunc("/api/creditcards/{id}", controllers.GetCreditCardByID).Methods("GET")
	router.HandleFunc("/api/creditcards", controllers.UpdateCreditCardByID).Methods("PUT")
	router.HandleFunc("/api/creditcards/{id}", controllers.DeleteCreditCardByID).Methods("DELETE")

	return router
}
