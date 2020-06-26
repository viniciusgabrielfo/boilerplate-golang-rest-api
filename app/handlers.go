package app

import (
	"boilerplate/controllers"

	"github.com/gorilla/mux"
)

func ChargeRoutes() *mux.Router {
	router := mux.NewRouter()

	// User handlers
	router.HandleFunc("/api/users", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/api/users", controllers.GetAllUsers).Methods("GET")
	router.HandleFunc("/api/users/{id}", controllers.GetUserByID).Methods("GET")
	router.HandleFunc("/api/users", controllers.UpdateUserByID).Methods("PUT")
	router.HandleFunc("/api/users/{id}", controllers.DeleteUserByID).Methods("DELETE")

	// Credit Card handlers
	router.HandleFunc("/api/creditcards", controllers.CreateCreditCard).Methods("POST")
	router.HandleFunc("/api/creditcards", controllers.GetAllCreditCards).Methods("GET")
	router.HandleFunc("/api/creditcards/{id}", controllers.GetCreditCardByID).Methods("GET")
	router.HandleFunc("/api/creditcards", controllers.UpdateCreditCardByID).Methods("PUT")
	router.HandleFunc("/api/creditcards/{id}", controllers.DeleteCreditCardByID).Methods("DELETE")

	return router
}
