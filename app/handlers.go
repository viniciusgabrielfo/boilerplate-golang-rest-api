package app

import (
	"boilerplate/controllers"

	"github.com/gorilla/mux"
)

func ChargeRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/users", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/api/users", controllers.GetAllUsers).Methods("GET")
	router.HandleFunc("/api/users/{id}", controllers.GetUserById).Methods("GET")
	router.HandleFunc("/api/users", controllers.UpdateUserById).Methods("PUT")
	router.HandleFunc("/api/users/{id}", controllers.DeleteUserById).Methods("DELETE")

	return router
}
