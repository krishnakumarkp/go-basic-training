package router

import (
	"go-training/customer/controller"
	"go-training/customer/domain"

	"github.com/gorilla/mux"
)

func GetRouter(customerStore domain.CustomerStore) *mux.Router {

	controller := controller.CustomerController{
		Store: customerStore,
	}

	router := mux.NewRouter()

	router.HandleFunc("/customer", controller.Add).Methods("POST")
	router.HandleFunc("/customer/{id}", controller.Get).Methods("GET")
	router.HandleFunc("/customer", controller.GetAll).Methods("GET")
	router.HandleFunc("/customer/{id}", controller.Delete).Methods("DELETE")
	router.HandleFunc("/customer/{id}", controller.Update).Methods("PUT")
	return router
}
