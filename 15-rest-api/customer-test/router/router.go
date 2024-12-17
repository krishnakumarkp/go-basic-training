package router

import (
	"go-training/customer-map/apistore"
	"go-training/customer-map/controller"
	"go-training/customer-map/domain"
	"net/http"
)

func GetRouter(customerStore domain.CustomerStore) *http.ServeMux {
	apiStore := apistore.OrderStore{BaseURL: "http://localhost:8082"}

	controller := controller.NewCustomerController(customerStore, apiStore)

	router := http.NewServeMux()

	router.HandleFunc("POST /customer", controller.Add)
	router.HandleFunc("GET /customer/{id}", controller.Get)
	router.HandleFunc("GET /customer/{id}/order", controller.GetOrders)

	return router
}
