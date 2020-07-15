package app

import (
	"boilerplate/controllers"
	u "boilerplate/utils"
	"net/http"

	"github.com/gorilla/mux"
)

// LoadRoutes is a function to create and return a new *mux.Router and your routes.
func LoadRoutes() *mux.Router {
	router := mux.NewRouter()

	// Auth handlers
	router.HandleFunc("/api/login", controllers.Signin).Methods("POST")
	router.HandleFunc("/api/refresh", controllers.Refresh).Methods("POST")

	// User handlers
	router.HandleFunc("/api/users", controllers.CreateUser).Methods("POST")
	router.Handle("/api/users", u.IsAuthorizedMiddleware(http.HandlerFunc(controllers.GetAllUsers))).Methods("GET")
	router.Handle("/api/users/{id}", u.IsAuthorizedMiddleware(http.HandlerFunc(controllers.GetUserByID))).Methods("GET")
	router.Handle("/api/users", u.IsAuthorizedMiddleware(http.HandlerFunc(controllers.UpdateUserByID))).Methods("PUT")
	router.Handle("/api/users/{id}", u.IsAuthorizedMiddleware(http.HandlerFunc(controllers.DeleteUserByID))).Methods("DELETE")
	router.Handle("/api/users/{id}/creditcards", u.IsAuthorizedMiddleware(http.HandlerFunc(controllers.GetAllCreditCardsByUserID))).Methods("GET")

	// Credit Card handlers
	router.Handle("/api/creditcards", u.IsAuthorizedMiddleware(http.HandlerFunc(controllers.CreateCreditCard))).Methods("POST")
	router.Handle("/api/creditcards", u.IsAuthorizedMiddleware(http.HandlerFunc(controllers.GetAllCreditCards))).Methods("GET")
	router.Handle("/api/creditcards/{id}", u.IsAuthorizedMiddleware(http.HandlerFunc(controllers.GetCreditCardByID))).Methods("GET")
	router.Handle("/api/creditcards", u.IsAuthorizedMiddleware(http.HandlerFunc(controllers.UpdateCreditCardByID))).Methods("PUT")
	router.Handle("/api/creditcards/{id}", u.IsAuthorizedMiddleware(http.HandlerFunc(controllers.DeleteCreditCardByID))).Methods("DELETE")

	return router
}
